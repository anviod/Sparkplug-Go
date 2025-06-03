package config

import (
	"github.com/spf13/viper"
)

type DeviceConfig struct {
	ID      string             `mapstructure:"id"`
	Type    string             `mapstructure:"type"`
	Metrics map[string]float64 `mapstructure:"metrics"`
}

type NodeConfig struct {
	ID      string         `mapstructure:"id"`
	Devices []DeviceConfig `mapstructure:"devices"`
}

type MQTTConfig struct {
	Broker     string `mapstructure:"broker"`
	ClientID   string `mapstructure:"client_id"`
	Username   string `mapstructure:"username"`
	Password   string `mapstructure:"password"`
	GroupID    string `mapstructure:"group_id"`
	LWTTopic   string `mapstructure:"lwt_topic"`
	LWTPayload string `mapstructure:"lwt_payload"`
	QoS        int    `mapstructure:"qos"`
}

type Config struct {
	MQTT MQTTConfig `mapstructure:"mqtt"`
	Node NodeConfig `mapstructure:"node"`
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
