//
// Copyright (c) 2017
// Cavium
//
// SPDX-License-Identifier: Apache-2.0
//

package client

const (
	defaultPort       = 48071
	defaultHostname   = "127.0.0.1"
	defaultDistroHost = "127.0.0.1"
	defaultConsulHost = "127.0.0.1"
	defaultConsulPort = 8500
)

type Config struct {
	Port       int
	Hostname   string
	DistroHost string
	ConsulHost string
	ConsulPort int
}

var cfg Config

func GetDefaultConfig() Config {
	return Config{
		Port:       defaultPort,
		Hostname:   defaultHostname,
		DistroHost: defaultDistroHost,
		ConsulHost: defaultConsulHost,
		ConsulPort: defaultConsulPort,
	}
}
