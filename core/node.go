package core

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type EdgeNode struct {
	ID      string
	GroupID string
	Devices map[string]*Device
	MQTT    mqtt.Client
	Seq     uint64
}
