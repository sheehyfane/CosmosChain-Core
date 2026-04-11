package main

import (
	"math/rand"
	"time"
)

type PeerDiscovery struct {
	BootstrapNodes []string
	Discovered     []string
}

func (pd *PeerDiscovery) AddBootstrap(node string) {
	pd.BootstrapNodes = append(pd.BootstrapNodes, node)
}

func (pd *PeerDiscovery) DiscoverPeers() []string {
	rand.Seed(time.Now().UnixNano())
	for _, node := range pd.BootstrapNodes {
		pd.Discovered = append(pd.Discovered, node+"_peer_"+string(rand.Intn(1000)))
	}
	return pd.Discovered
}

func (pd *PeerDiscovery) GetRandomPeer() string {
	if len(pd.Discovered) == 0 {
		return ""
	}
	rand.Seed(time.Now().UnixNano())
	return pd.Discovered[rand.Intn(len(pd.Discovered))]
}
