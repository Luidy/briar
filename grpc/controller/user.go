package controller

import (
	"context"
	"ges/model"

	"google.golang.org/grpc"
)

type GRPCUserHandler struct {
}

func newGRPCUserHandler(g *grpc.Server) {
	handler := &GRPCUserHandler{}
	model.RegisterUserServer(g, handler)
}

func (h *GRPCUserHandler) RegistUser(ctx context.Context, req *model.UserInfo) (res *model.UserInfo, err error) {
	return &model.UserInfo{}, nil
}
