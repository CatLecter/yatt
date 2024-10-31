package grpcclient

import (
	"context"
	"github.com/CatLecter/yatt/internal/dto/bff"
	userv1 "github.com/CatLecter/yatt/pkg/gen/go/v1/user"
	"github.com/google/uuid"
	grpclog "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type Client struct {
	api userv1.UserClient
	log *zerolog.Logger
}

func New(
	log *zerolog.Logger,
	addr string,
	retryTimeout time.Duration,
	retries uint,
) (*Client, error) {
	retryOpts := []grpcretry.CallOption{
		grpcretry.WithCodes(codes.NotFound, codes.Aborted, codes.DeadlineExceeded),
		grpcretry.WithMax(retries),
		grpcretry.WithPerRetryTimeout(retryTimeout),
	}

	logOpts := []grpclog.Option{
		grpclog.WithLogOnEvents(grpclog.PayloadReceived, grpclog.PayloadSent),
	}
	cc, err := grpc.NewClient(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(
			grpclog.UnaryClientInterceptor(InterceptorLogger(log), logOpts...),
			grpcretry.UnaryClientInterceptor(retryOpts...),
		),
	)
	if err != nil {
		return nil, err
	}
	return &Client{api: userv1.NewUserClient(cc), log: log}, nil
}

func InterceptorLogger(log *zerolog.Logger) grpclog.Logger {
	return grpclog.LoggerFunc(func(ctx context.Context, lvl grpclog.Level, msg string, fields ...any) {
		switch lvl {
		case grpclog.LevelDebug:
			log.Debug().Fields(fields).Msg(msg)
		case grpclog.LevelInfo:
			log.Info().Fields(fields).Msg(msg)
		case grpclog.LevelWarn:
			log.Warn().Fields(fields).Msg(msg)
		case grpclog.LevelError:
			log.Error().Fields(fields).Msg(msg)
		default:
			log.Panic().Msgf("unknown level %v", lvl)
		}
	})
}

func (c *Client) Create(ctx context.Context, req *dto.CreateUserRequest) (*dto.UserBriefResponse, error) {
	resp, err := c.api.Create(
		ctx,
		&userv1.CreateUserRequest{
			Username:        req.UserName,
			FullName:        req.FullName,
			Email:           req.Email,
			Password:        req.Password,
			ConfirmPassword: req.ConfirmPassword,
		},
	)
	if err != nil {
		// TODO: маппинг gRPC ошибок на HTTP.
		return nil, err
	}
	return &dto.UserBriefResponse{UserID: uuid.Must(uuid.Parse(resp.GetUserId()))}, nil
}

func (c *Client) Get(ctx context.Context, req *dto.UserBriefRequest) (*dto.UserResponse, error) {
	resp, err := c.api.Get(ctx, &userv1.UserBriefRequest{UserId: req.UserID.String()})
	if err != nil {
		// TODO: маппинг gRPC ошибок на HTTP.
		return nil, err
	}
	return &dto.UserResponse{
		UserID:       uuid.Must(uuid.Parse(resp.GetUserId())),
		UserName:     resp.GetUsername(),
		FullName:     resp.GetFullName(),
		Email:        resp.GetEmail(),
		Active:       resp.GetActive(),
		Hidden:       resp.GetHidden(),
		LastLogin:    resp.GetLastLogin().AsTime(),
		CustomFields: resp.GetCustomFields(),
		CreatedAt:    resp.GetCreatedAt().AsTime(),
		UpdatedAt:    resp.GetUpdatedAt().AsTime(),
	}, nil
}

func (c *Client) Update(ctx context.Context, req *dto.UpdateUserRequest) (*dto.UserResponse, error) {
	resp, err := c.api.Update(
		ctx,
		&userv1.UpdateUserRequest{
			UserId:       req.UserID.String(),
			Username:     req.UserName,
			FullName:     req.FullName,
			Email:        req.Email,
			CustomFields: req.CustomFields,
		},
	)
	if err != nil {
		// TODO: маппинг gRPC ошибок на HTTP.
		return nil, err
	}
	return &dto.UserResponse{
		UserID:       uuid.Must(uuid.Parse(resp.GetUserId())),
		UserName:     resp.GetUsername(),
		FullName:     resp.GetFullName(),
		Email:        resp.GetEmail(),
		Active:       resp.GetActive(),
		Hidden:       resp.GetHidden(),
		LastLogin:    resp.GetLastLogin().AsTime(),
		CustomFields: resp.GetCustomFields(),
		CreatedAt:    resp.GetCreatedAt().AsTime(),
		UpdatedAt:    resp.GetUpdatedAt().AsTime(),
	}, nil
}

func (c *Client) Login(ctx context.Context, req *dto.LoginRequest) (*dto.LoginResponse, error) {
	resp, err := c.api.Login(ctx, &userv1.LoginRequest{
		Username: req.UserName,
		Password: req.Password,
	})
	if err != nil {
		// TODO: маппинг gRPC ошибок на HTTP.
		return nil, err
	}
	return &dto.LoginResponse{AccessToken: resp.GetAccessToken()}, nil
}
