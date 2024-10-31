package service

import (
	"context"
	domainUser "github.com/CatLecter/yatt/internal/domain/user"
	repository "github.com/CatLecter/yatt/internal/infrastructure/repositories/iaa"
	"github.com/CatLecter/yatt/internal/lib"
	userv1 "github.com/CatLecter/yatt/pkg/gen/go/v1/user"
	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserService struct {
	log         *zerolog.Logger
	userStorage repository.UserStorageInterface
}

func NewUserService(
	log *zerolog.Logger,
	userStorage repository.UserStorageInterface,
) UserServiceInterface {
	return &UserService{log: log, userStorage: userStorage}
}

func (s *UserService) Create(ctx *context.Context, req *userv1.CreateUserRequest) (*userv1.UserBriefResponse, *lib.GRPCError) {
	user, err := domainUser.Register(
		req.GetUsername(), req.GetFullName(), req.GetEmail(), req.GetPassword(), req.GetConfirmPassword(),
	)
	if err != nil {
		return nil, err
	}
	if err := s.userStorage.Save(ctx, user); err != nil {
		return nil, lib.NewGRPCError(codes.Internal, err.Error())
	}
	return &userv1.UserBriefResponse{UserId: user.ID}, nil
}

func (s *UserService) Get(ctx *context.Context, req *userv1.UserBriefRequest) (*userv1.UserResponse, *lib.GRPCError) {
	user := domainUser.New()
	if err := user.SetID(req.GetUserId()); err != nil {
		return nil, err
	}
	if err := s.userStorage.GetByID(ctx, user); err != nil {
		return nil, lib.NewGRPCError(codes.NotFound, err.Error())
	}
	return &userv1.UserResponse{
		UserId:       user.ID,
		Username:     user.UserName,
		FullName:     user.FullName,
		Email:        user.Email,
		Active:       user.Active,
		Hidden:       user.Hidden,
		LastLogin:    timestamppb.New(user.LastLogin),
		CustomFields: user.CustomFields,
		CreatedAt:    timestamppb.New(user.CreatedAt),
		UpdatedAt:    timestamppb.New(user.UpdatedAt),
	}, nil
}

func (s *UserService) Update(ctx *context.Context, req *userv1.UpdateUserRequest) (*userv1.UserResponse, *lib.GRPCError) {
	user := domainUser.New()
	if err := user.SetID(req.GetUserId()); err != nil {
		return nil, err
	}
	if err := s.userStorage.GetByID(ctx, user); err != nil {
		return nil, lib.NewGRPCError(codes.NotFound, err.Error())
	}
	user.SelfUpdate(req.GetUsername(), req.GetFullName(), req.GetEmail(), req.GetCustomFields())
	if err := s.userStorage.Update(ctx, user); err != nil {
		return nil, lib.NewGRPCError(codes.Internal, err.Error())
	}
	return &userv1.UserResponse{
		UserId:       user.ID,
		Username:     user.UserName,
		FullName:     user.FullName,
		Email:        user.Email,
		Active:       user.Active,
		Hidden:       user.Hidden,
		LastLogin:    timestamppb.New(user.LastLogin),
		CustomFields: user.CustomFields,
		CreatedAt:    timestamppb.New(user.CreatedAt),
		UpdatedAt:    timestamppb.New(user.UpdatedAt),
	}, nil
}

func (s *UserService) Login(ctx *context.Context, req *userv1.LoginRequest) (*userv1.LoginResponse, *lib.GRPCError) {
	user := domainUser.New()
	if err := user.SetUserName(req.GetUsername()); err != nil {
		return nil, err
	}
	if err := s.userStorage.GetByUserName(ctx, user); err != nil {
		return nil, lib.NewGRPCError(codes.NotFound, err.Error())
	}
	accessToken, err := user.CreateAccessToken(req.GetPassword())
	if err != nil {
		return nil, err
	}
	return &userv1.LoginResponse{AccessToken: accessToken}, nil
}
