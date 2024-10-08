// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package zipkinudpreceiver

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/zipkinudpreceiver/internal/metadata"
)

func createDefaultConfig() component.Config {
	return &Config{
		ListenAddress: defaultListenAddress,
		UDPConfig:     defaultConfigUDP(),
	}
}

func createTracesReceiver(
	ctx context.Context,
	set receiver.Settings,
	cfg component.Config,
	nextConsumer consumer.Traces,
) (receiver.Traces, error) {
	zCfg := cfg.(*Config)

	return newZipkinUDPReceiver(ctx, &set, zCfg, nextConsumer)
}

func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		metadata.Type,
		createDefaultConfig,
		receiver.WithTraces(createTracesReceiver, metadata.TracesStability),
	)
}
