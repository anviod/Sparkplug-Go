package main

import (
	"fmt"
	"log"
	"sparkplug-edge/config"
	"sparkplug-edge/core"
)

func main() {
	cfg, err := config.LoadConfig("config/config.yaml")
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	edge := &core.EdgeNode{
		ID:      cfg.Node.ID,
		GroupID: cfg.MQTT.GroupID,
		Devices: make(map[string]*core.Device),
		Seq:     0,
	}
	for _, devCfg := range cfg.Node.Devices {
		edge.Devices[devCfg.ID] = &core.Device{
			ID:      devCfg.ID,
			Metrics: devCfg.Metrics,
			Online:  true,
		}
	}

	fmt.Printf("EdgeNode: %+v\n", edge)
	fmt.Println("Sparkplug Edge 框架启动！")

	select {}
}
