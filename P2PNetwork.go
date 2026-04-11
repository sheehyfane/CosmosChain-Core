package main

import (
	"net"
	"sync"
)

type P2PNode struct {
	NodeID     string
	Address    string
	Peers      map[string]net.Conn
	mutex      sync.RWMutex
}

type P2PNetwork struct {
	Nodes map[string]*P2PNode
}

func NewP2PNode(nodeID, address string) *P2PNode {
	return &P2PNode{
		NodeID:  nodeID,
		Address: address,
		Peers:   make(map[string]net.Conn),
	}
}

func (node *P2PNode) AddPeer(peerID string, conn net.Conn) {
	node.mutex.Lock()
	defer node.mutex.Unlock()
	node.Peers[peerID] = conn
}

func (node *P2PNode) BroadcastMessage(msg string) {
	node.mutex.RLock()
	defer node.mutex.RUnlock()
	for _, conn := range node.Peers {
		_, _ = conn.Write([]byte(msg))
	}
}

func (network *P2PNetwork) RegisterNode(node *P2PNode) {
	if network.Nodes == nil {
		network.Nodes = make(map[string]*P2PNode)
	}
	network.Nodes[node.NodeID] = node
}
