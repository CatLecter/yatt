package identitygrpc

import (
	"context"
	"github.com/CatLecter/yatt/internal/schemes"
	userv1 "github.com/CatLecter/yatt/pkg/gen/go/v1/user"
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
	timeout time.Duration,
	retriesCount uint,
) (*Client, error) {
	retryOpts := []grpcretry.CallOption{
		grpcretry.WithCodes(codes.NotFound, codes.Aborted, codes.DeadlineExceeded),
		grpcretry.WithMax(retriesCount),
		grpcretry.WithPerRetryTimeout(timeout),
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

func (c *Client) Get(ctx context.Context, req *schemes.UserBriefRequest) (*schemes.UserResponse, error) {
	resp, err := c.api.Get(ctx, &userv1.UserBriefRequest{UserId: req.UserID})
	if err != nil {
		return nil, err
	}
	return &schemes.UserResponse{
		UserID:    resp.GetUserId(),
		UserName:  resp.GetUsername(),
		Email:     resp.GetEmail(),
		CreatedAt: resp.GetCreatedAt().AsTime(),
		UpdatedAt: resp.GetUpdatedAt().AsTime(),
	}, nil
}
func (c *Client) Create(ctx context.Context, req *schemes.UserRequest) (*schemes.UserBriefResponse, error) {
	resp, err := c.api.Create(
		ctx,
		&userv1.UserRequest{
			Username:        req.UserName,
			Email:           req.Email,
			Password:        req.Password,
			ConfirmPassword: req.ConfirmPassword,
		},
	)
	if err != nil {
		return nil, err
	}
	return &schemes.UserBriefResponse{UserID: resp.GetUserId()}, nil
}

func (c *Client) Update(ctx context.Context, req *schemes.UserRequest) (*schemes.UserResponse, error) {
	resp, err := c.api.Update(
		ctx,
		&userv1.UserRequest{
			Username:        req.UserName,
			Email:           req.Email,
			Password:        req.Password,
			ConfirmPassword: req.ConfirmPassword,
		},
	)
	if err != nil {
		return nil, err
	}
	return &schemes.UserResponse{
		UserID:    resp.GetUserId(),
		UserName:  resp.GetUsername(),
		Email:     resp.GetEmail(),
		CreatedAt: resp.GetCreatedAt().AsTime(),
		UpdatedAt: resp.GetUpdatedAt().AsTime(),
	}, nil
}

func (c *Client) Delete(ctx context.Context, req *schemes.UserBriefRequest) (*schemes.UserBriefResponse, error) {
	resp, err := c.api.Get(ctx, &userv1.UserBriefRequest{UserId: req.UserID})
	if err != nil {
		return nil, err
	}
	return &schemes.UserBriefResponse{UserID: resp.GetUserId()}, nil
}
