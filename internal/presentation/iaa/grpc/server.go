package server

import (
	"google.golang.org/grpc"
	"yatt/internal/services/iaa"
	userv1 "yatt/pkg/gen/go/v1/user"
)

type userServer struct {
	userv1.UnimplementedUserServer
	service service.UserServiceInterface
}

func New(gRPC *grpc.Server, service service.UserServiceInterface) {
	userv1.RegisterUserServer(gRPC, &userServer{service: service})
}
