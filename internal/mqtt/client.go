package mqtt

import (
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// MQTT 客户端封装
type Client struct {
	client mqtt.Client
}

// Config 用于 MQTT 客户端配置
type Config struct {
	Broker     string // MQTT 代理地址
	ClientID   string // 客户端 ID
	Username   string // 用户名
	Password   string // 密码
	LWTTopic   string // 最后遗嘱主题
	LWTPayload string // 最后遗嘱负载
	QoS        int    // 服务质量
}

// NewClient 创建一个新的 MQTT 客户端实例
func NewClient(cfg Config, onMessage mqtt.MessageHandler) (*Client, error) {
	opts := mqtt.NewClientOptions().AddBroker(cfg.Broker)
	opts.SetClientID(cfg.ClientID)
	opts.SetUsername(cfg.Username)
	opts.SetPassword(cfg.Password)
	opts.SetKeepAlive(30 * time.Second)
	opts.SetPingTimeout(10 * time.Second)
	opts.SetAutoReconnect(true)
	opts.SetConnectRetry(true)
	opts.SetConnectRetryInterval(5 * time.Second)
	opts.SetDefaultPublishHandler(onMessage)
	opts.SetWill(cfg.LWTTopic, cfg.LWTPayload, byte(cfg.QoS), false)

	cli := mqtt.NewClient(opts)
	if token := cli.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}
	return &Client{client: cli}, nil
}

// Publish 发布消息到指定主题
func (c *Client) Publish(topic string, qos byte, retained bool, payload []byte) error {
	token := c.client.Publish(topic, qos, retained, payload)
	token.Wait()
	return token.Error()
}

// Subscribe 订阅指定主题
func (c *Client) Subscribe(topic string, qos byte, callback mqtt.MessageHandler) error {
	token := c.client.Subscribe(topic, qos, callback)
	token.Wait()
	return token.Error()
}

// Disconnect 断开与 MQTT 代理的连接
func (c *Client) Disconnect(quiesce uint) {
	c.client.Disconnect(quiesce)
}
