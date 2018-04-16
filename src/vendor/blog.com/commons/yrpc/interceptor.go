package yrpc

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	// github扩展库：https://github.com/grpc-ecosystem/go-grpc-middleware
)

// ServerInterceptor : grpc.UnaryServerInterceptor
func ServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	resp, err = handler(ctx, req)

	if md, ok := metadata.FromIncomingContext(ctx); !ok {
		fmt.Println("Server interceptor get metadata failed.")
	} else {
		fmt.Println("[测] Server interceptor metadata:", md)
	}

	return
}

// ClientInterceptor : grpc.UnaryClientInterceptor
func ClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	// md, ok := metadata.FromOutgoingContext(ctx)
	err := invoker(ctx, method, req, reply, cc, opts...)
	return err
}
