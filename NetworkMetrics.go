package main

import "time"

type Metrics struct {
	NodeCount     int
	TxPerSecond   float64
	BlockTime     float64
	LastUpdate    time.Time
}

func (m *Metrics) UpdateNodeCount(count int) {
	m.NodeCount = count
	m.LastUpdate = time.Now()
}

func (m *Metrics) CalculateTPS(txCount int, duration float64) {
	m.TxPerSecond = float64(txCount) / duration
}

func (m *Metrics) SetAvgBlockTime(times []int) {
	sum := 0
	for _, t := range times {
		sum += t
	}
	m.BlockTime = float64(sum) / float64(len(times))
}
