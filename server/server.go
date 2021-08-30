package server

import (
	"briar/config"
	"briar/idl"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type UserServer struct {
	cfg config.Config
}

func NewUserServer(cfg config.Config) (*UserServer, error) {
	return &UserServer{cfg: cfg}, nil
}

func (u *UserServer) Echo(ctx context.Context, req *idl.EchoRequest) (*idl.EchoResponse, error) {
	return nil, nil
}

func NewServer(ctx context.Context, cfg config.Config) (*grpc.Server, error) {
	grpcServer := grpc.NewServer()
	userServer, err := NewUserServer(cfg)
	if err != nil {
		return nil, err
	}

	idl.RegisterUserServer(grpcServer, userServer)
	reflection.Register(grpcServer)

	return grpcServer, nil
}
