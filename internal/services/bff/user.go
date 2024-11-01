package service

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"net/http"
	dto "yatt/internal/dto/bff"
	grpcclient "yatt/internal/infrastructure/clients/bff/grpc"
	"yatt/internal/lib"
	userv1 "yatt/pkg/gen/go/v1/user"
)

type UserService struct {
	log        *zerolog.Logger
	userClient *grpcclient.Client
}

func NewUserService(log *zerolog.Logger, userClient *grpcclient.Client) UserServiceInterface {
	return &UserService{log: log, userClient: userClient}
}

func (s *UserService) Create(ctx *gin.Context) {
	req := dto.CreateUserRequest{}
	if err := ctx.BindJSON(&req); err != nil {
		s.log.Warn().Err(err).Msg("cannot parse body")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, lib.NewHTTPError("cannot parse body"))
		return
	}
	resp, err := s.userClient.Create(ctx, &userv1.CreateUserRequest{
		Username:        req.UserName,
		FullName:        req.FullName,
		Email:           req.Email,
		Password:        req.Password,
		ConfirmPassword: req.ConfirmPassword,
	})
	if err != nil {
		// TODO: реализовать маппинг статус-кодов gRPC на HTTP.
		s.log.Warn().Err(err).Msg("user creation error")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, lib.NewHTTPError("user creation error"))
		return
	}
	ctx.JSON(http.StatusOK, dto.UserBriefResponse{UserID: uuid.MustParse(resp.GetUserId())})
	return
}

func (s *UserService) Get(ctx *gin.Context) {
	resp, err := s.userClient.Get(ctx, &userv1.UserBriefRequest{UserId: ctx.Param("user_id")})
	if err != nil {
		// TODO: реализовать маппинг статус-кодов gRPC на HTTP.
		s.log.Warn().Err(err).Msg("user not found")
		ctx.AbortWithStatusJSON(http.StatusNotFound, lib.NewHTTPError("user not found"))
		return
	}
	ctx.JSON(http.StatusOK, dto.UserResponse{
		UserID:       uuid.MustParse(resp.GetUserId()),
		UserName:     resp.GetUsername(),
		FullName:     resp.GetFullName(),
		Email:        resp.GetEmail(),
		Active:       resp.GetActive(),
		Hidden:       resp.GetHidden(),
		LastLogin:    resp.GetLastLogin().AsTime(),
		CustomFields: resp.GetCustomFields(),
		CreatedAt:    resp.GetCreatedAt().AsTime(),
		UpdatedAt:    resp.GetUpdatedAt().AsTime(),
	})
	return
}

func (s *UserService) Update(ctx *gin.Context) {
	req := dto.UpdateUserRequest{}
	if err := ctx.BindJSON(&req); err != nil {
		s.log.Warn().Err(err).Msg("cannot parse body")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, lib.NewHTTPError("cannot parse body"))
		return
	}
	resp, err := s.userClient.Update(ctx, &userv1.UpdateUserRequest{
		UserId:       ctx.Param("user_id"),
		Username:     req.UserName,
		FullName:     req.FullName,
		Email:        req.Email,
		CustomFields: req.CustomFields,
	})
	if err != nil {
		// TODO: реализовать маппинг статус-кодов gRPC на HTTP.
		s.log.Warn().Err(err).Msg("user updating error")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, lib.NewHTTPError("user updating error"))
		return
	}
	ctx.JSON(http.StatusOK, dto.UserResponse{
		UserID:       uuid.MustParse(resp.GetUserId()),
		UserName:     resp.GetUsername(),
		FullName:     resp.GetFullName(),
		Email:        resp.GetEmail(),
		Active:       resp.GetActive(),
		Hidden:       resp.GetHidden(),
		LastLogin:    resp.GetLastLogin().AsTime(),
		CustomFields: resp.GetCustomFields(),
		CreatedAt:    resp.GetCreatedAt().AsTime(),
		UpdatedAt:    resp.GetUpdatedAt().AsTime(),
	})
	return
}

func (s *UserService) Login(ctx *gin.Context) {
	req := dto.LoginRequest{}
	if err := ctx.BindJSON(&req); err != nil {
		s.log.Warn().Err(err).Msg("cannot parse body")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, lib.NewHTTPError("cannot parse body"))
		return
	}
	resp, err := s.userClient.Login(ctx, &userv1.LoginRequest{
		Username: req.UserName,
		Password: req.Password,
	})
	if err != nil {
		// TODO: реализовать маппинг статус-кодов gRPC на HTTP.
		s.log.Warn().Err(err).Msg("user login error")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, lib.NewHTTPError("user login error"))
		return
	}
	ctx.JSON(http.StatusOK, dto.LoginResponse{AccessToken: resp.GetAccessToken()})
	return
}
