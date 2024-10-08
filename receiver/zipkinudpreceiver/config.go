// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package zipkinudpreceiver

import (
	"fmt"
	"net"
	"strconv"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/zipkinudpreceiver/internal/metadata"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap"
)

type Config struct {
	ListenAddress        string                        `mapstructure:"listen_address"`
	UDPConfig            ProtocolUDP                   `mapstructure:"udp"`
	MetricsBuilderConfig metadata.MetricsBuilderConfig `mapstructure:",squash"`
}

var _ component.Config = (*Config)(nil)
var _ confmap.Unmarshaler = (*Config)(nil)

const (
	udpFieldName = "udp"

	defaultListenAddress = "localhost:5775"
	defaultMaxPacketSize = 65_000
	defaultNumWorkers    = 10
	defaultQueueSize     = 1000
)

func (cfg *Config) Validate() error {
	if cfg.ListenAddress == "" {
		return fmt.Errorf("listen address is empty")
	}

	if err := checkPortFromAddress(cfg.ListenAddress); err != nil {
		return err
	}

	if cfg.UDPConfig.QueueSize <= 0 {
		return fmt.Errorf("queue size must be positive")
	}

	if cfg.UDPConfig.NumWorkers <= 0 {
		return fmt.Errorf("number of workers must be positive")
	}

	if cfg.UDPConfig.MaxPacketSize <= 2 {
		return fmt.Errorf("max packet size must be at least 3")
	}

	return nil
}

func (cfg *Config) Unmarshal(componentParser *confmap.Conf) error {
	err := componentParser.Unmarshal(cfg)
	if err != nil {
		return err
	}
	listenAddress := cfg.ListenAddress
	if listenAddress == "" {
		return fmt.Errorf("listen address is empty")
	}
	if err := checkPortFromAddress(listenAddress); err != nil {
		return err
	}
	udp, err := componentParser.Sub(udpFieldName)
	if err != nil {
		return err
	}
	if udp != nil {
		err = udp.Unmarshal(&cfg.UDPConfig)
		if err != nil {
			return err
		}
	}

	return nil
}

type ProtocolUDP struct {
	ServerConfigUDP `mapstructure:",squash"`
}

type ServerConfigUDP struct {
	QueueSize     int `mapstructure:"queue_size"`
	NumWorkers    int `mapstructure:"num_workers"`
	MaxPacketSize int `mapstructure:"max_packet_size"`
}

func defaultConfigUDP() ProtocolUDP {
	return ProtocolUDP{
		ServerConfigUDP: ServerConfigUDP{
			QueueSize:     defaultQueueSize,
			NumWorkers:    defaultNumWorkers,
			MaxPacketSize: defaultMaxPacketSize,
		},
	}
}

func defaultConfig() *Config {
	return &Config{
		ListenAddress: defaultListenAddress,
		UDPConfig:     defaultConfigUDP(),
	}
}

func checkPortFromAddress(address string) error {
	// Check if the port is valid
	_, portStr, err := net.SplitHostPort(address)
	if err != nil {
		return fmt.Errorf("invalid address: %w", err)
	}
	port, err := strconv.ParseInt(portStr, 10, 0)
	if err != nil {
		return fmt.Errorf("invalid port, must be a number: %w", err)
	}
	if port < 0 || port > 65535 {
		return fmt.Errorf("invalid port, must be between 0 and 65535: %d", port)
	}
	return nil
}
