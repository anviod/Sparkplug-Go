package core

type Device struct {
	ID      string
	Metrics map[string]float64
	Online  bool
}
