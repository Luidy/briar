package controller

import "google.golang.org/grpc"

func InitHandler(g *grpc.Server) (err error) {
	newGRPCUserHandler(g)
	return nil
}
