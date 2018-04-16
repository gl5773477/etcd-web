package yrpc

import (
	"context"
	"sync"

	"git.ymt360.com/go/gocommons/logging"
)

var yrpcCtxPool = sync.Pool{
	New: func() interface{} {
		return new(YRPCContext)
	},
}

type YRPCContext struct {
	ctx       context.Context
	LogHeader *logging.LogHeader
}

func NewYRPCContext() *YRPCContext {
	yctx := yrpcCtxPool.Get().(*YRPCContext)

	return yctx
}
