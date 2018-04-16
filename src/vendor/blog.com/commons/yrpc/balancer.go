package yrpc

import (
	"git.ymt360.com/go/gocommons/yrpc/naming"
	"google.golang.org/grpc"
)

// TODO
// remote grpclb.

// RoundRobin default.
func RoundRobin(r *naming.YRPCResolver) grpc.Balancer {
	return grpc.RoundRobin(r)
}
