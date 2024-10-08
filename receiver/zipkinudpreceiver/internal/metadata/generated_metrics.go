// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
)

type metricReceiverZipkinUDPCountBufferFull struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills receiver_zipkin_udp_count_buffer_full metric with initial data.
func (m *metricReceiverZipkinUDPCountBufferFull) init() {
	m.data.SetName("receiver_zipkin_udp_count_buffer_full")
	m.data.SetDescription("Count of the number of times the internal buffer was full and we dropped spans.")
	m.data.SetUnit("{events}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricReceiverZipkinUDPCountBufferFull) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricReceiverZipkinUDPCountBufferFull) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricReceiverZipkinUDPCountBufferFull) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricReceiverZipkinUDPCountBufferFull(cfg MetricConfig) metricReceiverZipkinUDPCountBufferFull {
	m := metricReceiverZipkinUDPCountBufferFull{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricReceiverZipkinUDPCountSpansDropped struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills receiver_zipkin_udp_count_spans_dropped metric with initial data.
func (m *metricReceiverZipkinUDPCountSpansDropped) init() {
	m.data.SetName("receiver_zipkin_udp_count_spans_dropped")
	m.data.SetDescription("Count of spans dropped by the Zipkin UDP receiver.")
	m.data.SetUnit("{spans}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricReceiverZipkinUDPCountSpansDropped) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricReceiverZipkinUDPCountSpansDropped) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricReceiverZipkinUDPCountSpansDropped) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricReceiverZipkinUDPCountSpansDropped(cfg MetricConfig) metricReceiverZipkinUDPCountSpansDropped {
	m := metricReceiverZipkinUDPCountSpansDropped{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricReceiverZipkinUDPCountSpansFailedToParse struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills receiver_zipkin_udp_count_spans_failed_to_parse metric with initial data.
func (m *metricReceiverZipkinUDPCountSpansFailedToParse) init() {
	m.data.SetName("receiver_zipkin_udp_count_spans_failed_to_parse")
	m.data.SetDescription("Count of how often we fail to decode the incoming spans.")
	m.data.SetUnit("{spans}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricReceiverZipkinUDPCountSpansFailedToParse) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricReceiverZipkinUDPCountSpansFailedToParse) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricReceiverZipkinUDPCountSpansFailedToParse) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricReceiverZipkinUDPCountSpansFailedToParse(cfg MetricConfig) metricReceiverZipkinUDPCountSpansFailedToParse {
	m := metricReceiverZipkinUDPCountSpansFailedToParse{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricReceiverZipkinUDPCountSpansReceived struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills receiver_zipkin_udp_count_spans_received metric with initial data.
func (m *metricReceiverZipkinUDPCountSpansReceived) init() {
	m.data.SetName("receiver_zipkin_udp_count_spans_received")
	m.data.SetDescription("Count of spans received by the Zipkin UDP receiver.")
	m.data.SetUnit("{spans}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricReceiverZipkinUDPCountSpansReceived) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricReceiverZipkinUDPCountSpansReceived) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricReceiverZipkinUDPCountSpansReceived) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricReceiverZipkinUDPCountSpansReceived(cfg MetricConfig) metricReceiverZipkinUDPCountSpansReceived {
	m := metricReceiverZipkinUDPCountSpansReceived{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricReceiverZipkinUDPCountUDPOverflows struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills receiver_zipkin_udp_count_udp_overflows metric with initial data.
func (m *metricReceiverZipkinUDPCountUDPOverflows) init() {
	m.data.SetName("receiver_zipkin_udp_count_udp_overflows")
	m.data.SetDescription("Count of how often we have received a too big UDP package.")
	m.data.SetUnit("{events}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricReceiverZipkinUDPCountUDPOverflows) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricReceiverZipkinUDPCountUDPOverflows) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricReceiverZipkinUDPCountUDPOverflows) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricReceiverZipkinUDPCountUDPOverflows(cfg MetricConfig) metricReceiverZipkinUDPCountUDPOverflows {
	m := metricReceiverZipkinUDPCountUDPOverflows{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user config.
type MetricsBuilder struct {
	config                                         MetricsBuilderConfig // config of the metrics builder.
	startTime                                      pcommon.Timestamp    // start time that will be applied to all recorded data points.
	metricsCapacity                                int                  // maximum observed number of metrics per resource.
	metricsBuffer                                  pmetric.Metrics      // accumulates metrics data before emitting.
	buildInfo                                      component.BuildInfo  // contains version information.
	metricReceiverZipkinUDPCountBufferFull         metricReceiverZipkinUDPCountBufferFull
	metricReceiverZipkinUDPCountSpansDropped       metricReceiverZipkinUDPCountSpansDropped
	metricReceiverZipkinUDPCountSpansFailedToParse metricReceiverZipkinUDPCountSpansFailedToParse
	metricReceiverZipkinUDPCountSpansReceived      metricReceiverZipkinUDPCountSpansReceived
	metricReceiverZipkinUDPCountUDPOverflows       metricReceiverZipkinUDPCountUDPOverflows
}

// metricBuilderOption applies changes to default metrics builder.
type metricBuilderOption func(*MetricsBuilder)

// WithStartTime sets startTime on the metrics builder.
func WithStartTime(startTime pcommon.Timestamp) metricBuilderOption {
	return func(mb *MetricsBuilder) {
		mb.startTime = startTime
	}
}

func NewMetricsBuilder(mbc MetricsBuilderConfig, settings receiver.Settings, options ...metricBuilderOption) *MetricsBuilder {
	mb := &MetricsBuilder{
		config:                                   mbc,
		startTime:                                pcommon.NewTimestampFromTime(time.Now()),
		metricsBuffer:                            pmetric.NewMetrics(),
		buildInfo:                                settings.BuildInfo,
		metricReceiverZipkinUDPCountBufferFull:   newMetricReceiverZipkinUDPCountBufferFull(mbc.Metrics.ReceiverZipkinUDPCountBufferFull),
		metricReceiverZipkinUDPCountSpansDropped: newMetricReceiverZipkinUDPCountSpansDropped(mbc.Metrics.ReceiverZipkinUDPCountSpansDropped),
		metricReceiverZipkinUDPCountSpansFailedToParse: newMetricReceiverZipkinUDPCountSpansFailedToParse(mbc.Metrics.ReceiverZipkinUDPCountSpansFailedToParse),
		metricReceiverZipkinUDPCountSpansReceived:      newMetricReceiverZipkinUDPCountSpansReceived(mbc.Metrics.ReceiverZipkinUDPCountSpansReceived),
		metricReceiverZipkinUDPCountUDPOverflows:       newMetricReceiverZipkinUDPCountUDPOverflows(mbc.Metrics.ReceiverZipkinUDPCountUDPOverflows),
	}
	for _, op := range options {
		op(mb)
	}
	return mb
}

// updateCapacity updates max length of metrics and resource attributes that will be used for the slice capacity.
func (mb *MetricsBuilder) updateCapacity(rm pmetric.ResourceMetrics) {
	if mb.metricsCapacity < rm.ScopeMetrics().At(0).Metrics().Len() {
		mb.metricsCapacity = rm.ScopeMetrics().At(0).Metrics().Len()
	}
}

// ResourceMetricsOption applies changes to provided resource metrics.
type ResourceMetricsOption func(pmetric.ResourceMetrics)

// WithResource sets the provided resource on the emitted ResourceMetrics.
// It's recommended to use ResourceBuilder to create the resource.
func WithResource(res pcommon.Resource) ResourceMetricsOption {
	return func(rm pmetric.ResourceMetrics) {
		res.CopyTo(rm.Resource())
	}
}

// WithStartTimeOverride overrides start time for all the resource metrics data points.
// This option should be only used if different start time has to be set on metrics coming from different resources.
func WithStartTimeOverride(start pcommon.Timestamp) ResourceMetricsOption {
	return func(rm pmetric.ResourceMetrics) {
		var dps pmetric.NumberDataPointSlice
		metrics := rm.ScopeMetrics().At(0).Metrics()
		for i := 0; i < metrics.Len(); i++ {
			switch metrics.At(i).Type() {
			case pmetric.MetricTypeGauge:
				dps = metrics.At(i).Gauge().DataPoints()
			case pmetric.MetricTypeSum:
				dps = metrics.At(i).Sum().DataPoints()
			}
			for j := 0; j < dps.Len(); j++ {
				dps.At(j).SetStartTimestamp(start)
			}
		}
	}
}

// EmitForResource saves all the generated metrics under a new resource and updates the internal state to be ready for
// recording another set of data points as part of another resource. This function can be helpful when one scraper
// needs to emit metrics from several resources. Otherwise calling this function is not required,
// just `Emit` function can be called instead.
// Resource attributes should be provided as ResourceMetricsOption arguments.
func (mb *MetricsBuilder) EmitForResource(rmo ...ResourceMetricsOption) {
	rm := pmetric.NewResourceMetrics()
	ils := rm.ScopeMetrics().AppendEmpty()
	ils.Scope().SetName("otelcol/zipkinudpreceiver")
	ils.Scope().SetVersion(mb.buildInfo.Version)
	ils.Metrics().EnsureCapacity(mb.metricsCapacity)
	mb.metricReceiverZipkinUDPCountBufferFull.emit(ils.Metrics())
	mb.metricReceiverZipkinUDPCountSpansDropped.emit(ils.Metrics())
	mb.metricReceiverZipkinUDPCountSpansFailedToParse.emit(ils.Metrics())
	mb.metricReceiverZipkinUDPCountSpansReceived.emit(ils.Metrics())
	mb.metricReceiverZipkinUDPCountUDPOverflows.emit(ils.Metrics())

	for _, op := range rmo {
		op(rm)
	}
	if ils.Metrics().Len() > 0 {
		mb.updateCapacity(rm)
		rm.MoveTo(mb.metricsBuffer.ResourceMetrics().AppendEmpty())
	}
}

// Emit returns all the metrics accumulated by the metrics builder and updates the internal state to be ready for
// recording another set of metrics. This function will be responsible for applying all the transformations required to
// produce metric representation defined in metadata and user config, e.g. delta or cumulative.
func (mb *MetricsBuilder) Emit(rmo ...ResourceMetricsOption) pmetric.Metrics {
	mb.EmitForResource(rmo...)
	metrics := mb.metricsBuffer
	mb.metricsBuffer = pmetric.NewMetrics()
	return metrics
}

// RecordReceiverZipkinUDPCountBufferFullDataPoint adds a data point to receiver_zipkin_udp_count_buffer_full metric.
func (mb *MetricsBuilder) RecordReceiverZipkinUDPCountBufferFullDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricReceiverZipkinUDPCountBufferFull.recordDataPoint(mb.startTime, ts, val)
}

// RecordReceiverZipkinUDPCountSpansDroppedDataPoint adds a data point to receiver_zipkin_udp_count_spans_dropped metric.
func (mb *MetricsBuilder) RecordReceiverZipkinUDPCountSpansDroppedDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricReceiverZipkinUDPCountSpansDropped.recordDataPoint(mb.startTime, ts, val)
}

// RecordReceiverZipkinUDPCountSpansFailedToParseDataPoint adds a data point to receiver_zipkin_udp_count_spans_failed_to_parse metric.
func (mb *MetricsBuilder) RecordReceiverZipkinUDPCountSpansFailedToParseDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricReceiverZipkinUDPCountSpansFailedToParse.recordDataPoint(mb.startTime, ts, val)
}

// RecordReceiverZipkinUDPCountSpansReceivedDataPoint adds a data point to receiver_zipkin_udp_count_spans_received metric.
func (mb *MetricsBuilder) RecordReceiverZipkinUDPCountSpansReceivedDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricReceiverZipkinUDPCountSpansReceived.recordDataPoint(mb.startTime, ts, val)
}

// RecordReceiverZipkinUDPCountUDPOverflowsDataPoint adds a data point to receiver_zipkin_udp_count_udp_overflows metric.
func (mb *MetricsBuilder) RecordReceiverZipkinUDPCountUDPOverflowsDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricReceiverZipkinUDPCountUDPOverflows.recordDataPoint(mb.startTime, ts, val)
}

// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...metricBuilderOption) {
	mb.startTime = pcommon.NewTimestampFromTime(time.Now())
	for _, op := range options {
		op(mb)
	}
}
