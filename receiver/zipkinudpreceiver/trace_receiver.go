// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package zipkinudpreceiver

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"sync"
	"time"

	zipkinmodel "github.com/openzipkin/zipkin-go/model"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/receiverhelper"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/zipkin/zipkinv2"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/zipkinudpreceiver/internal/metadata"
)

type zipkinUDPReceiver struct {
	ctx             context.Context
	cancel          context.CancelFunc
	nextTraces      consumer.Traces
	config          *Config
	settings        *receiver.Settings
	obsrepUDP       *receiverhelper.ObsReport
	mb              *metadata.MetricsBuilder
	jsonUnmarshaler ptrace.Unmarshaler
	shutdownWG      sync.WaitGroup
	queue           chan []byte
	translator      zipkinv2.ToTranslator
	logger          *zap.Logger
}

const (
	udpTransport   = "udp"
	zipkinV2Format = "zipkinv2"
)

func (z *zipkinUDPReceiver) Start(_ context.Context, host component.Host) error {
	if err := z.startUDPServer(host); err != nil {
		return err
	}
	return nil
}

// startUDPServer starts the UDP server and the workers that process the incoming messages.
func (z *zipkinUDPReceiver) startUDPServer(host component.Host) error {
	if host == nil {
		return fmt.Errorf("host is nil")
	}
	var err error
	addr, err := net.ResolveUDPAddr("udp", z.config.ListenAddress)
	if err != nil {
		return fmt.Errorf("failed to resolve UDP address: %w", err)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return fmt.Errorf("failed to listen on UDP address %s: %w", z.config.ListenAddress, err)
	}
	z.logger.Debug("Started Zipkin UDP Receiver",
		zap.String("address", z.config.ListenAddress),
		zap.Int("num_workers", z.config.UDPConfig.NumWorkers),
		zap.Int("queue_size", z.config.UDPConfig.QueueSize),
		zap.Int("max_packet_size", z.config.UDPConfig.MaxPacketSize))

	messages := make(chan []byte, z.config.UDPConfig.QueueSize)
	for i := 0; i < z.config.UDPConfig.NumWorkers; i++ {
		z.shutdownWG.Add(1)
		go z.worker(i, messages)
	}
	ctx, cancel := context.WithCancel(context.Background())
	z.cancel = cancel
	go z.listenUDP(ctx, conn, messages)
	return nil
}

// worker consumes messages from the channel and processes them.
func (z *zipkinUDPReceiver) worker(id int, messages <-chan []byte) {
	defer z.shutdownWG.Done()
	now := pcommon.NewTimestampFromTime(time.Now())
	for message := range messages {
		z.logger.Debug("Worker received message",
			zap.Int("worker_id", id),
			zap.String("message", string(message)))
		ctx := z.obsrepUDP.StartTracesOp(context.Background())
		var spans []*zipkinmodel.SpanModel
		err := json.Unmarshal(message, &spans)
		if err != nil {
			z.obsrepUDP.EndTracesOp(ctx, "json.Unmarshal", 0, err)
			z.logger.Error("Error unmarshalling Zipkin V2 JSON", zap.Error(err))
			z.mb.RecordReceiverZipkinUDPCountSpansFailedToParseDataPoint(now, 1)
			continue
		}
		td, err := z.translator.ToTraces(spans)
		if err != nil {
			z.logger.Error("Error translating spans to traces", zap.Error(err))
			z.mb.RecordReceiverZipkinUDPCountSpansFailedToParseDataPoint(now, 1)
			z.obsrepUDP.EndTracesOp(ctx, zipkinV2Format, 0, err)
			continue
		}
		numReceivedSpans := td.SpanCount()
		consumerErr := z.nextTraces.ConsumeTraces(ctx, td)
		z.obsrepUDP.EndTracesOp(ctx, zipkinV2Format, numReceivedSpans, consumerErr)
	}
}

// listenUDP listens on the given UDP connection and sends incoming messages to the given channel.
func (z *zipkinUDPReceiver) listenUDP(
	ctx context.Context,
	conn *net.UDPConn,
	messages chan<- []byte,
) {
	defer z.shutdownWG.Done()
	buffer := make([]byte, z.config.UDPConfig.MaxPacketSize)
	now := pcommon.NewTimestampFromTime(time.Now())
	for {
		select {
		case <-ctx.Done():
			return
		default:
			n, _, err := conn.ReadFromUDP(buffer)
			if n == 0 {
				continue
			}
			if n > z.config.UDPConfig.MaxPacketSize {
				z.logger.Warn("Dropping message because it exceeds the maximum packet size",
					zap.Int("actual_size", n),
					zap.Int("max_size", z.config.UDPConfig.MaxPacketSize))
				z.mb.RecordReceiverZipkinUDPCountUDPOverflowsDataPoint(now, 1)
				continue
			}
			if err != nil {
				z.logger.Error("Error reading from UDP connection", zap.Error(err))
				continue
			}
			data := make([]byte, n)
			copy(data, buffer[:n])

			select {
			case messages <- data:
			default:
				z.logger.Warn("Dropping message because the queue is full",
					zap.Int("queue_size", len(messages)),
					zap.Int("max_queue_size", cap(messages)))
				z.mb.RecordReceiverZipkinUDPCountBufferFullDataPoint(now, 1)
			}
		}
	}
}

func (z *zipkinUDPReceiver) Shutdown(_ context.Context) error {
	if z.cancel != nil {
		z.cancel()
	}
	z.shutdownWG.Wait()
	close(z.queue)
	return nil
}

func newZipkinUDPReceiver(
	ctx context.Context,
	settings *receiver.Settings,
	cfg *Config,
	nextConsumer consumer.Traces,
) (*zipkinUDPReceiver, error) {
	telemetrySettings := settings.TelemetrySettings
	udpObsrecv, err := receiverhelper.NewObsReport(receiverhelper.ObsReportSettings{
		ReceiverID:             settings.ID,
		Transport:              udpTransport,
		ReceiverCreateSettings: *settings,
	})
	if err != nil {
		return nil, err
	}
	return &zipkinUDPReceiver{
		ctx:             ctx,
		nextTraces:      nextConsumer,
		config:          cfg,
		settings:        settings,
		mb:              metadata.NewMetricsBuilder(cfg.MetricsBuilderConfig, *settings),
		obsrepUDP:       udpObsrecv,
		logger:          telemetrySettings.Logger,
		queue:           make(chan []byte, cfg.UDPConfig.QueueSize),
		jsonUnmarshaler: zipkinv2.NewJSONTracesUnmarshaler(true),
	}, nil
}
