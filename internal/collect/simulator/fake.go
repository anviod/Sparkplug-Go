package simulator

type FakeCollector struct{}

func NewFakeCollector() *FakeCollector {
	return &FakeCollector{}
}

func (f *FakeCollector) Init(config map[string]interface{}) error {
	return nil
}

func (f *FakeCollector) Collect() (map[string]float64, error) {
	return map[string]float64{"temp": 25.0, "humidity": 60.0}, nil
}
