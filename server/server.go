package server

import (
	"briar/config"
	"briar/idl"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type BriarServer struct {
	cfg config.Config
}

func initBriarServer(cfg config.Config) *BriarServer {
	return &BriarServer{
		cfg: cfg,
	}
}

var _ idl.BriarServer = (*BriarServer)(nil)

func (s *BriarServer) Echo(ctx context.Context, req *idl.EchoRequest) (*idl.EchoResponse, error) {
	return nil, nil
}

func InitGRPCServer(cfg config.Config) (*grpc.Server, error) {
	grpcServer := grpc.NewServer()
	briarServer := initBriarServer(cfg)

	idl.RegisterBriarServer(grpcServer, briarServer)
	reflection.Register(grpcServer)

	return grpcServer, nil
}
