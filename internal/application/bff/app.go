package bff

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"strings"
	"time"
	docs "yatt/api/v1/swagger"
	"yatt/internal/presentation/bff/handlers"
)

type App struct {
	server *http.Server
	log    *zerolog.Logger
}

func New(
	handler *handlers.Handler,
	log *zerolog.Logger,
	host string,
	port int,
	readTimeout time.Duration,
	writeTimeout time.Duration,
	allowOrigins string,
) *App {

	engine := gin.New()

	engine.Use(
		cors.New(
			cors.Config{
				AllowOrigins:     strings.Split(allowOrigins, ";"),
				AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
				AllowHeaders:     []string{"Origin", "Content-type", "Authorization"},
				AllowCredentials: true,
			},
		),
	)

	engine.Use(
		gin.LoggerWithFormatter(
			func(param gin.LogFormatterParams) string {
				log.Debug().
					Str("client_ip", param.ClientIP).
					Str("method", param.Method).
					Str("path", param.Path).
					Int("status_code", param.StatusCode).
					Str("latency", fmt.Sprintf("%v us", param.Latency.Microseconds())).
					Msg(param.ErrorMessage)
				return ""
			},
		),
	)

	docs.SwaggerInfo.Title = "Swagger API Docs - OpenAPI 2.0"
	docs.SwaggerInfo.Description = "API documentation for the Gin template"

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := engine.Group("/v1")

	auth := v1.Group("/auth")
	{
		auth.POST("/login", handler.Login)

	}

	user := v1.Group("/user")
	{
		user.POST("/", handler.Create)
		user.GET("/:user_id", handler.Get)
		user.PUT("/:user_id", handler.Update)

		return &App{
			log: log,
			server: &http.Server{
				Addr:           fmt.Sprintf("%s:%v", host, port),
				Handler:        engine,
				MaxHeaderBytes: http.DefaultMaxHeaderBytes,
				ReadTimeout:    readTimeout,
				WriteTimeout:   writeTimeout,
			},
		}
	}
}

func (a *App) Run() error {
	a.log.Info().Msg("Server is running: " + a.server.Addr)
	if err := a.server.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
}

func (a *App) Stop(ctx context.Context) error {
	a.log.Info().Msg("Stopping server...")
	return a.server.Shutdown(ctx)
}
