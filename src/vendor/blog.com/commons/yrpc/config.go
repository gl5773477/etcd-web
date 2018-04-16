package yrpc

import (
	"time"
)

type YRPCClientConfig struct {
	EtcdCfg *EtcdConfig
}

type YRPCServerConfig struct {
	EtcdCfg *EtcdConfig
	Name    string
	Host    string
	TTL     int64
}

type EtcdConfig struct {
	Addr        string
	Endpoints   []string
	DialTimeout time.Duration
}
