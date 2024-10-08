// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package zipkinudpreceiver

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/pipeline"
	"go.opentelemetry.io/collector/receiver/receivertest"
)

func TestTypeStr(t *testing.T) {
	factory := NewFactory()
	require.Equal(t, "zipkin_udp", factory.Type().String())
}

func TestCreateDefaultConfig(t *testing.T) {
	factory := NewFactory()
	cfg := factory.CreateDefaultConfig()
	require.NotNil(t, cfg, "failed to create default config")
	assert.NoError(t, componenttest.CheckConfigStruct(cfg))
}

func TestCreateReceiver(t *testing.T) {
	factory := NewFactory()
	cfg := factory.CreateDefaultConfig()
	set := receivertest.NewNopSettings()
	tReceiver, err := factory.CreateTracesReceiver(context.Background(), set, cfg, nil)
	assert.NoError(t, err, "receiver creation failed")
	assert.NotNil(t, tReceiver, "receiver creation failed")

	mReceiver, err := factory.CreateMetricsReceiver(context.Background(), set, cfg, nil)
	assert.Equal(t, err, pipeline.ErrSignalNotSupported)
	assert.Nil(t, mReceiver)

	lReceiver, err := factory.CreateLogsReceiver(context.Background(), set, cfg, nil)
	assert.Equal(t, err, pipeline.ErrSignalNotSupported)
	assert.Nil(t, lReceiver)
}

func TestCreateTraces(t *testing.T) {
	factory := NewFactory()
	cfg := factory.CreateDefaultConfig()
	set := receivertest.NewNopSettings()
	r, err := factory.CreateTracesReceiver(context.Background(), set, cfg, nil)

	assert.NoError(t, err, "unexpected error creating receiver")
	assert.Equal(t, "localhost:5775", r.(*zipkinUDPReceiver).config.ListenAddress)
}
