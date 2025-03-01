package sidecar

import "time"

type Sidecar struct {
	Alias    string    `yaml:"alias"`
	Address  string    `yaml:"address"`
	Status   string    `yaml:"status"`
	LastPing time.Time `yaml:"last_ping"`
}
