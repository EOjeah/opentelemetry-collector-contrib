// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package zipkinudpreceiver

import "go.opentelemetry.io/collector/component"

// import (
// 	"go.opentelemetry.io/collector/component"
// 	"go.opentelemetry.io/collector/receiver/receivertest"
// )

var zipkinUDPReceiverID = component.MustNewIDWithName("zipkin_udp", "receiver_test")

// // func TestTraceSource(t *testing.T) {
// // 	set := receivertest.NewNopSettings()
// // 	zr, err := newZipkinUDPReceiver(zipkinudpReceiver, &c)
// // }
