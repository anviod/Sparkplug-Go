package proto

import (
	"sparkplug-edge/internal/proto/sparkplug_b"
	"time"

	"google.golang.org/protobuf/proto"
)

// Sparkplug Protobuf 报文封装占位
type MetricType int

const (
	MetricDouble MetricType = iota
	MetricInt
	MetricString
)

type Metric struct {
	Name  string
	Type  MetricType
	Value interface{}
}

type PayloadBuilder struct {
	Metrics   []Metric
	Seq       uint64
	Timestamp int64
}

func NewPayloadBuilder(seq uint64) *PayloadBuilder {
	return &PayloadBuilder{
		Metrics:   []Metric{},
		Seq:       seq,
		Timestamp: time.Now().UnixMilli(),
	}
}

func (b *PayloadBuilder) AddMetric(name string, typ MetricType, value interface{}) {
	b.Metrics = append(b.Metrics, Metric{Name: name, Type: typ, Value: value})
}

func (b *PayloadBuilder) Build() *sparkplug_b.Payload {
	metrics := []*sparkplug_b.Metric{}
	for _, m := range b.Metrics {
		metric := &sparkplug_b.Metric{
			Name: m.Name,
		}
		switch m.Type {
		case MetricDouble:
			metric.Datatype = 0
			if v, ok := m.Value.(float64); ok {
				metric.DoubleValue = v
			}
		case MetricInt:
			metric.Datatype = 1
			if v, ok := m.Value.(int64); ok {
				metric.IntValue = v
			}
		case MetricString:
			metric.Datatype = 2
			if v, ok := m.Value.(string); ok {
				metric.StringValue = v
			}
		}
		metrics = append(metrics, metric)
	}
	return &sparkplug_b.Payload{
		Timestamp: b.Timestamp,
		Seq:       b.Seq,
		Metrics:   metrics,
	}
}

func MarshalPayload(payload *sparkplug_b.Payload) ([]byte, error) {
	return proto.Marshal(payload)
}
