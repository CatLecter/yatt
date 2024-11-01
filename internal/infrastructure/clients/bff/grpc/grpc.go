package grpcclient

import (
	"context"
	grpclog "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"time"
	userv1 "yatt/pkg/gen/go/v1/user"
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
		grpcretry.WithCodes(codes.Aborted, codes.DeadlineExceeded),
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

func (c *Client) Create(ctx context.Context, req *userv1.CreateUserRequest) (*userv1.UserBriefResponse, error) {
	return c.api.Create(ctx, req)
}

func (c *Client) Get(ctx context.Context, req *userv1.UserBriefRequest) (*userv1.UserResponse, error) {
	return c.api.Get(ctx, req)
}

func (c *Client) Update(ctx context.Context, req *userv1.UpdateUserRequest) (*userv1.UserResponse, error) {
	return c.api.Update(ctx, req)
}

func (c *Client) Login(ctx context.Context, req *userv1.LoginRequest) (*userv1.LoginResponse, error) {
	return c.api.Login(ctx, req)
}
