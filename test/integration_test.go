package test

import (
	"sparkplug-edge/internal/app"
	"sparkplug-edge/internal/collect/simulator"
	"sparkplug-edge/internal/model"
	"testing"
)

func TestDeviceDataPush(t *testing.T) {
	node := app.NewEdgeNode("edge-1", "group-1")
	dev := &model.Device{
		ID:        "dev-001",
		Collector: simulator.NewFakeCollector(),
		Metrics:   map[string]float64{},
		Online:    true,
	}
	node.RegisterDevice(dev)

	// TODO: 启动采集与数据推送，验证 MQTT 是否收到 DBIRTH/DDATA
}
