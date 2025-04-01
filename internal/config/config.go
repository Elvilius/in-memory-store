package config

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Engine  *Engine  `yaml:"engine"`
	Network *Network `yaml:"network"`
	Logging *Logging `yaml:"logging"`
}

type Engine struct {
	Type             string `yaml:"type"`
	PartitionsNumber uint   `yaml:"partitions_number"`
}

type Network struct {
	Address        string        `yaml:"address"`
	MaxConnections int           `yaml:"max_connections"`
	MaxMessageSize string        `yaml:"max_message_size"`
	IdleTimeout    time.Duration `yaml:"idle_timeout"`
	BufferSize     int           `yaml:"buffer_size"`
}

type Logging struct {
	Level  string `yaml:"level"`
	Output string `yaml:"output"`
}

func New() (*Config, error) {
	data, err := os.ReadFile(getConfigPath())
	if err != nil {
		return nil, err
	}

	var cfg Config

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	if cfg.Network.MaxConnections == 0 {
		return nil, fmt.Errorf("cfg.Network.MaxConnections must be > 0")
	}

	if cfg.Network.BufferSize == 0 {
		cfg.Network.BufferSize = 4 << 10 // 4096
	}
	return &cfg, nil
}

func getConfigPath() string {
	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		path = "./config.yaml"
	}
	return path
}
