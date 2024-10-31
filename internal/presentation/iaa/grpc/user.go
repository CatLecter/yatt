package server

import (
	"context"
	userv1 "github.com/CatLecter/yatt/pkg/gen/go/v1/user"
	"google.golang.org/grpc/status"
)

func (srv *userServer) Create(ctx context.Context, req *userv1.CreateUserRequest) (*userv1.UserBriefResponse, error) {
	resp, err := srv.service.Create(&ctx, req)
	if err != nil {
		return nil, status.Error(err.Code, err.Msg)
	}
	return resp, nil
}

func (srv *userServer) Get(ctx context.Context, req *userv1.UserBriefRequest) (*userv1.UserResponse, error) {
	resp, err := srv.service.Get(&ctx, req)
	if err != nil {
		return nil, status.Error(err.Code, err.Msg)
	}
	return resp, nil
}

func (srv *userServer) Update(ctx context.Context, req *userv1.UpdateUserRequest) (*userv1.UserResponse, error) {
	resp, err := srv.service.Update(&ctx, req)
	if err != nil {
		return nil, status.Error(err.Code, err.Msg)
	}
	return resp, nil
}

func (srv *userServer) Login(ctx context.Context, req *userv1.LoginRequest) (*userv1.LoginResponse, error) {
	resp, err := srv.service.Login(&ctx, req)
	if err != nil {
		return nil, status.Error(err.Code, err.Msg)
	}
	return resp, nil
}
