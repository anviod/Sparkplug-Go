package model

import "sparkplug-edge/internal/collect"

type Device struct {
	ID        string
	Collector collect.Collector
	Metrics   map[string]float64
	Online    bool
}
