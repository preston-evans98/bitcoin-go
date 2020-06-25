package config

import (
	"bitcoin-go/pkg/constants"
	"sync"
)

// NodeConfig is an object representing a Node's configuration at runtime
type NodeConfig struct {
	network int
	mu      *sync.Mutex
}

func (conf *NodeConfig) GetNetwork() int {
	conf.mu.Lock()
	defer conf.mu.Unlock()
	return conf.network
}

func (conf *NodeConfig) SetNetwork(network int) {
	conf.mu.Lock()
	defer conf.mu.Unlock()
	conf.network = network
}

func GenerateDefaultConfig() *NodeConfig {
	conf := &NodeConfig{}
	conf.network = constants.Mainnet
	conf.mu = &sync.Mutex{}
	return conf
}
