package server

import (
	"github.com/CatLecter/yatt/internal/services/iaa"
	userv1 "github.com/CatLecter/yatt/pkg/gen/go/v1/user"
	"google.golang.org/grpc"
)

type userServer struct {
	userv1.UnimplementedUserServer
	service service.UserServiceInterface
}

func New(gRPC *grpc.Server, service service.UserServiceInterface) {
	userv1.RegisterUserServer(gRPC, &userServer{service: service})
}
