package modbus

import "errors"

type ModbusCollector struct{}

func NewModbusCollector() *ModbusCollector {
	return &ModbusCollector{}
}

func (m *ModbusCollector) Init(config map[string]interface{}) error {
	// TODO: 初始化 Modbus TCP 客户端
	return errors.New("Modbus 初始化未实现")
}

func (m *ModbusCollector) Collect() (map[string]float64, error) {
	// TODO: 采集 Modbus 数据
	return nil, errors.New("Modbus 采集未实现")
}
