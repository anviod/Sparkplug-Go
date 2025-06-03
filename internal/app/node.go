package app

import (
	"sparkplug-edge/internal/model"
)

type Node struct {
	Edge *model.EdgeNode
}

func NewEdgeNode(id, groupID string) *Node {
	return &Node{
		Edge: &model.EdgeNode{
			ID:      id,
			GroupID: groupID,
			Devices: make(map[string]*model.Device),
			Seq:     0,
		},
	}
}

func (n *Node) RegisterDevice(dev *model.Device) {
	n.Edge.Devices[dev.ID] = dev
}
