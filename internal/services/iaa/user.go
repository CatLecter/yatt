package service

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/timestamppb"
	domainUser "yatt/internal/domain/user"
	repository "yatt/internal/infrastructure/repositories/iaa"
	"yatt/internal/lib"
	userv1 "yatt/pkg/gen/go/v1/user"
)

type UserService struct {
	log         *zerolog.Logger
	userStorage repository.UserStorageInterface
	db          *pgxpool.Pool
}

func NewUserService(
	log *zerolog.Logger,
	userStorage repository.UserStorageInterface,
	db *pgxpool.Pool,
) UserServiceInterface {
	return &UserService{log: log, userStorage: userStorage, db: db}
}

func (s *UserService) Create(ctx *context.Context, req *userv1.CreateUserRequest) (*userv1.UserBriefResponse, *lib.GRPCError) {
	// TODO: создание транзакции вынести в обертку.
	conn, err := s.db.Acquire(*ctx)
	if err != nil {
		return nil, lib.NewGRPCError(codes.Internal, err.Error())
	}
	defer conn.Release()

	tx, err := conn.Begin(*ctx)
	if err != nil {
		return nil, lib.NewGRPCError(codes.Internal, err.Error())
	}
	defer func() {
		if err != nil {
			if err := tx.Rollback(*ctx); err != nil {
				s.log.Warn().Err(err).Msg("failed to rollback transaction")
			}
		} else {
			if err = tx.Commit(*ctx); err != nil {
				s.log.Warn().Err(err).Msg("failed to commit transaction")
			}
		}
	}()

	user, grpcErr := domainUser.Register(
		req.GetUsername(), req.GetFullName(), req.GetEmail(), req.GetPassword(), req.GetConfirmPassword(),
	)
	if grpcErr != nil {
		return nil, grpcErr
	}
	if err := s.userStorage.Save(ctx, tx, user); err != nil {
		return nil, lib.NewGRPCError(codes.Internal, err.Error())
	}
	return &userv1.UserBriefResponse{UserId: user.ID}, nil
}

func (s *UserService) Get(ctx *context.Context, req *userv1.UserBriefRequest) (*userv1.UserResponse, *lib.GRPCError) {
	// TODO: создание транзакции вынести в обертку.
	conn, err := s.db.Acquire(*ctx)
	if err != nil {
		return nil, lib.NewGRPCError(codes.Internal, err.Error())
	}
	defer conn.Release()

	tx, err := conn.Begin(*ctx)
	if err != nil {
		return nil, lib.NewGRPCError(codes.Internal, err.Error())
	}
	defer func() {
		if err != nil {
			if err := tx.Rollback(*ctx); err != nil {
				s.log.Warn().Err(err).Msg("failed to rollback transaction")
			}
		} else {
			if err = tx.Commit(*ctx); err != nil {
				s.log.Warn().Err(err).Msg("failed to commit transaction")
			}
		}
	}()

	user := domainUser.New()
	if err := user.SetID(req.GetUserId()); err != nil {
		return nil, err
	}
	if err := s.userStorage.GetByID(ctx, tx, user); err != nil {
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
	// TODO: создание транзакции вынести в обертку.
	conn, err := s.db.Acquire(*ctx)
	if err != nil {
		return nil, lib.NewGRPCError(codes.Internal, err.Error())
	}
	defer conn.Release()

	tx, err := conn.Begin(*ctx)
	if err != nil {
		return nil, lib.NewGRPCError(codes.Internal, err.Error())
	}
	defer func() {
		if err != nil {
			if err := tx.Rollback(*ctx); err != nil {
				s.log.Warn().Err(err).Msg("failed to rollback transaction")
			}
		} else {
			if err = tx.Commit(*ctx); err != nil {
				s.log.Warn().Err(err).Msg("failed to commit transaction")
			}
		}
	}()

	user := domainUser.New()
	if err := user.SetID(req.GetUserId()); err != nil {
		return nil, err
	}
	if err := s.userStorage.GetByID(ctx, tx, user); err != nil {
		return nil, lib.NewGRPCError(codes.NotFound, err.Error())
	}
	user.SelfUpdate(req.GetUsername(), req.GetFullName(), req.GetEmail(), req.GetCustomFields())
	if err := s.userStorage.Update(ctx, tx, user); err != nil {
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
	// TODO: создание транзакции вынести в обертку.
	conn, err := s.db.Acquire(*ctx)
	if err != nil {
		return nil, lib.NewGRPCError(codes.Internal, err.Error())
	}
	defer conn.Release()

	tx, err := conn.Begin(*ctx)
	if err != nil {
		return nil, lib.NewGRPCError(codes.Internal, err.Error())
	}
	defer func() {
		if err != nil {
			if err := tx.Rollback(*ctx); err != nil {
				s.log.Warn().Err(err).Msg("failed to rollback transaction")
			}
		} else {
			if err = tx.Commit(*ctx); err != nil {
				s.log.Warn().Err(err).Msg("failed to commit transaction")
			}
		}
	}()

	user := domainUser.New()
	if err := user.SetUserName(req.GetUsername()); err != nil {
		return nil, err
	}
	if err := s.userStorage.GetByUserName(ctx, tx, user); err != nil {
		return nil, lib.NewGRPCError(codes.NotFound, err.Error())
	}
	accessToken, grpcErr := user.CreateAccessToken(req.GetPassword())
	if grpcErr != nil {
		return nil, grpcErr
	}
	// TODO: реализовать обновление поля last_login пользователя
	return &userv1.LoginResponse{AccessToken: accessToken}, nil
}
