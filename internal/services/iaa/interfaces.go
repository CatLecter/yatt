package service

import (
	"context"
	"github.com/CatLecter/yatt/internal/lib"
	userv1 "github.com/CatLecter/yatt/pkg/gen/go/v1/user"
)

type UserServiceInterface interface {
	Create(ctx *context.Context, req *userv1.CreateUserRequest) (*userv1.UserBriefResponse, *lib.GRPCError)
	Get(ctx *context.Context, req *userv1.UserBriefRequest) (*userv1.UserResponse, *lib.GRPCError)
	Update(ctx *context.Context, req *userv1.UpdateUserRequest) (*userv1.UserResponse, *lib.GRPCError)
	Login(ctx *context.Context, req *userv1.LoginRequest) (*userv1.LoginResponse, *lib.GRPCError)
}
