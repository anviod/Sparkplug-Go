package collect

type Collector interface {
	Init(config map[string]interface{}) error
	Collect() (map[string]float64, error)
}
