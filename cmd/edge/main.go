package main

import (
	"sparkplug-edge/internal/app"
	"sparkplug-edge/internal/collect/simulator"
	"sparkplug-edge/internal/model"
	"sparkplug-edge/internal/utils"
)

func main() {
	utils.Logger.Info("Sparkplug Edge 框架启动！")

	node := app.NewEdgeNode("edge-1", "group-1")
	dev := &model.Device{
		ID:        "dev-001",
		Collector: simulator.NewFakeCollector(),
		Metrics:   map[string]float64{},
		Online:    true,
	}
	node.RegisterDevice(dev)

	// TODO: 初始化 MQTT、采集与推送、生命周期管理
}
