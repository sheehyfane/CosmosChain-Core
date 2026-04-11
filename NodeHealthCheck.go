package main

import "time"

type HealthCheck struct {
	node *ChainNode
}

func (hc *HealthCheck) CheckConnection() bool {
	return hc.node.Status == "active"
}

func (hc *HealthCheck) CheckSync() bool {
	return hc.node.IsSync
}

func (hc *HealthCheck) LastSeen() string {
	return time.Now().String()
}
