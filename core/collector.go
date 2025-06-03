package core

// Collector 采集协议插件接口

type Collector interface {
	Init(config map[string]interface{}) error
	Collect() (map[string]float64, error)
	Close() error
}
