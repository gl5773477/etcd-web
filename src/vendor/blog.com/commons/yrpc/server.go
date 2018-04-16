package yrpc

import (
	"context"
	"fmt"
	"net"

	"git.ymt360.com/go/gocommons/yrpc/naming"
	"github.com/coreos/etcd/clientv3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type YRPCService struct {
	ID      string
	Host    string
	r       *naming.YRPCResolver
	cli     *clientv3.Client
	serv    *grpc.Server
	leaseID clientv3.LeaseID
	stopCh  chan error
}

type ServiceRegistFunc func(s *grpc.Server)

func NewService(cfg *YRPCServerConfig, opts ...grpc.ServerOption) *YRPCService {
	s := &YRPCService{
		ID:     cfg.Name,
		Host:   cfg.Host,
		stopCh: make(chan error),
	}

	s.cli = NewEtcdConn(cfg.EtcdCfg)
	s.r = naming.NewResolver(s.ID, s.cli)
	leaseRes, err := s.cli.Grant(context.TODO(), 5)
	if err == nil {
		s.leaseID = leaseRes.ID
	}

	opts = append(opts, grpc.UnaryInterceptor(ServerInterceptor))
	s.serv = grpc.NewServer(opts...)
	return s
}

func (s *YRPCService) Regist(registFunc ServiceRegistFunc) {
	registFunc(s.serv)
	reflection.Register(s.serv)
	s.r.Regist(context.TODO(), s.Host, clientv3.WithLease(s.leaseID))
}

func (s *YRPCService) Unregist() {
	s.r.Delete(context.TODO(), s.Host)
}

// Only once called by main.
func (s *YRPCService) Start() {
	lis, err := net.Listen("tcp", s.Host)
	if err != nil {
		fmt.Printf("%s on host %s listen error:%v", s.ID, s.Host, err)
		panic(err)
	}

	go s.keepAlive()
	if err := s.serv.Serve(lis); err != nil {
		panic(err)
	}
}

func (s *YRPCService) Stop() {
	s.stopCh <- nil
}

func (s *YRPCService) keepAlive() {
	ch, err := s.cli.KeepAlive(context.TODO(), s.leaseID)
	if err != nil {
		fmt.Printf("[%s] set keepAlive failed:%v", s.ID, err)
		return
	}

	revokeFunc := func() {
		_, err := s.cli.Revoke(context.TODO(), s.leaseID)
		if err != nil {
			fmt.Printf("revoke failed:%v\n", err)
			return
		}
		fmt.Printf("[%s] revoke success\n", s.ID)
	}

	for {
		select {
		case <-s.stopCh:
			revokeFunc()
			return
		case <-s.cli.Ctx().Done():
			fmt.Println("server closed.")
			return
		case lka, ok := <-ch:
			if !ok {
				fmt.Println("keep alive closed.")
				revokeFunc()
				return
			}
			fmt.Printf("service normal,ttl:%d\n", lka.TTL)
		}
	}
}
