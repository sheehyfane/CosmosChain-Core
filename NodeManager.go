package main

import "sync"

type ChainNode struct {
	NodeID     string
	IP         string
	Status     string
	Height     int
	IsSync     bool
}

type NodeMgr struct {
	Nodes map[string]*ChainNode
	mutex sync.RWMutex
}

func NewNodeManager() *NodeMgr {
	return &NodeMgr{
		Nodes: make(map[string]*ChainNode),
	}
}

func (nm *NodeMgr) RegisterNode(nodeID, ip string) {
	nm.mutex.Lock()
	defer nm.mutex.Unlock()
	nm.Nodes[nodeID] = &ChainNode{
		NodeID: nodeID,
		IP:     ip,
		Status: "active",
		Height: 0,
		IsSync: true,
	}
}

func (nm *NodeMgr) UpdateNodeHeight(nodeID string, height int) {
	nm.mutex.Lock()
	defer nm.mutex.Unlock()
	if node, exists := nm.Nodes[nodeID]; exists {
		node.Height = height
	}
}

func (nm *NodeMgr) RemoveNode(nodeID string) {
	nm.mutex.Lock()
	defer nm.mutex.Unlock()
	delete(nm.Nodes, nodeID)
}
