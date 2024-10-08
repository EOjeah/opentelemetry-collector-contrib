// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package zipkinudpreceiver

import (
	"path/filepath"
	"testing"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/zipkinudpreceiver/internal/metadata"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap/confmaptest"
)

func TestLoadConfig(t *testing.T) {
	t.Parallel()

	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)

	tests := []struct {
		id       component.ID
		expected component.Config
	}{
		{
			id:       component.NewIDWithName(metadata.Type, "defaults"),
			expected: defaultConfig(),
		},
		{
			id: component.NewIDWithName(metadata.Type, "defaultudp"),
			expected: &Config{
				ListenAddress: "localhost:5875",
				UDPConfig:     defaultConfigUDP(),
			},
		},
		{
			id: component.NewIDWithName(metadata.Type, "custom"),
			expected: &Config{
				ListenAddress: "localhost:5776",
				UDPConfig: ProtocolUDP{
					ServerConfigUDP: ServerConfigUDP{
						MaxPacketSize: 1_000,
						NumWorkers:    5,
						QueueSize:     50,
					},
				},
			},
		},
		{
			id: component.NewIDWithName(metadata.Type, "udponly"),
			expected: &Config{
				ListenAddress: defaultListenAddress,
				UDPConfig: ProtocolUDP{
					ServerConfigUDP: ServerConfigUDP{
						MaxPacketSize: 10,
						NumWorkers:    1,
						QueueSize:     80,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.id.String(), func(t *testing.T) {
			factory := NewFactory()
			cfg := factory.CreateDefaultConfig()

			sub, err := cm.Sub(tt.id.String())
			require.NoError(t, err)
			require.NoError(t, sub.Unmarshal((cfg)))

			assert.NoError(t, component.ValidateConfig(cfg))
			assert.Equal(t, tt.expected, cfg)
		})
	}
}

func TestFailedEmptyListenAddress(t *testing.T) {
	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)
	factory := NewFactory()
	cfg := factory.CreateDefaultConfig()
	sub, err := cm.Sub(component.NewIDWithName(metadata.Type, "empty_listen_address").String())
	require.NoError(t, err)
	err = sub.Unmarshal(cfg)
	assert.ErrorContains(t, err, "listen address is empty")
}

func TestFailedInvalidListenAddress(t *testing.T) {
	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)
	factory := NewFactory()
	cfg := factory.CreateDefaultConfig()
	sub, err := cm.Sub(component.NewIDWithName(metadata.Type, "invalid_listen_address").String())
	require.NoError(t, err)
	err = sub.Unmarshal(cfg)
	assert.ErrorContains(t, err, "invalid address")
}

func TestFailedInvalidPortNumber(t *testing.T) {
	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)
	factory := NewFactory()
	cfg := factory.CreateDefaultConfig()
	sub, err := cm.Sub(component.NewIDWithName(metadata.Type, "invalid_port_string").String())
	require.NoError(t, err)
	err = sub.Unmarshal(cfg)
	assert.ErrorContains(t, err, "invalid port, must be a number")
	sub, err = cm.Sub(component.NewIDWithName(metadata.Type, "negative_port").String())
	require.NoError(t, err)
	err = sub.Unmarshal(cfg)
	assert.ErrorContains(t, err, "invalid port, must be between 0 and 65535")
	sub, err = cm.Sub(component.NewIDWithName(metadata.Type, "invalid_port_range").String())
	require.NoError(t, err)
	err = sub.Unmarshal(cfg)
	assert.ErrorContains(t, err, "invalid port, must be between 0 and 65535")
}
