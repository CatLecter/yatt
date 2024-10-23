package user

import (
	"context"
	"fmt"
	userv1 "github.com/CatLecter/yatt/pkg/gen/go/v1/user"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type userServer struct {
	userv1.UnimplementedUserServer
}

func NewUserServer(gRPC *grpc.Server) {
	userv1.RegisterUserServer(gRPC, &userServer{})
}

func (s *userServer) Register(ctx context.Context, req *userv1.UserRequest) (*userv1.LoginResponse, error) {
	fmt.Println("Register", req.String())
	return &userv1.LoginResponse{AccessToken: req.String()}, nil
}

func (s *userServer) Login(ctx context.Context, req *userv1.LoginRequest) (*userv1.LoginResponse, error) {
	fmt.Println("Login", req.String())
	return &userv1.LoginResponse{AccessToken: req.String()}, nil
}

func (s *userServer) Get(ctx context.Context, req *userv1.UserBriefRequest) (*userv1.UserResponse, error) {
	fmt.Println("Get", req.String())
	return &userv1.UserResponse{
		UserId:    req.GetUserId(),
		Username:  "someone",
		Email:     "some@email.com",
		CreatedAt: timestamppb.Now(),
		UpdatedAt: timestamppb.Now(),
	}, nil
}

func (s *userServer) Create(ctx context.Context, req *userv1.UserRequest) (*userv1.UserBriefResponse, error) {
	fmt.Println("Create", req.String())
	return &userv1.UserBriefResponse{UserId: "1"}, nil
}

func (s *userServer) Update(ctx context.Context, req *userv1.UserRequest) (*userv1.UserResponse, error) {
	fmt.Println("Update", req.String())
	return &userv1.UserResponse{
		UserId:    "1",
		Username:  req.GetUsername(),
		Email:     req.GetEmail(),
		CreatedAt: timestamppb.Now(),
		UpdatedAt: timestamppb.Now(),
	}, nil
}

func (s *userServer) Delete(ctx context.Context, req *userv1.UserBriefRequest) (*userv1.UserBriefResponse, error) {
	fmt.Println("Delete", req.String())
	return &userv1.UserBriefResponse{UserId: "1"}, nil
}
