package model

type EdgeNode struct {
	ID      string
	GroupID string
	Devices map[string]*Device
	Seq     uint64
}
