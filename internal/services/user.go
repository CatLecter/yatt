package services

import (
	"fmt"
	"github.com/CatLecter/yatt/internal/common"
	identitygrpc "github.com/CatLecter/yatt/internal/infrastructure/clients/identity/grpc"
	"github.com/CatLecter/yatt/internal/schemes"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"net/http"
)

type UserService struct {
	client *identitygrpc.Client
	log    *zerolog.Logger
}

func NewUserService(client *identitygrpc.Client, log *zerolog.Logger) UserServiceInterface {
	return &UserService{client: client, log: log}
}

func (srv *UserService) Get(ctx *gin.Context) {
	userID, err := uuid.Parse(ctx.Param("user_id"))
	if err != nil {
		srv.log.Err(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, common.NewHTTPError("Cannot parse user_id"))
		return
	}
	resp, err := srv.client.Get(ctx, &schemes.UserBriefRequest{UserID: userID.String()})
	if err != nil {
		srv.log.Err(err)
		ctx.AbortWithStatusJSON(
			http.StatusNotFound,
			common.NewHTTPError(fmt.Sprintf("User with ID %s not found", userID.String())),
		)
		return
	}
	ctx.JSON(http.StatusOK, resp)
	return
}

func (srv *UserService) Create(ctx *gin.Context) {
	req := schemes.UserRequest{}
	if err := ctx.BindJSON(&req); err != nil {
		srv.log.Err(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, common.NewHTTPError(fmt.Sprintf("Cannot parse JSON")))
		return
	}
	resp, err := srv.client.Create(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusUnprocessableEntity,
			common.NewHTTPError(fmt.Sprintf("identity with phone number %s already exists", req.UserName)),
		)
		return
	}
	ctx.JSON(http.StatusOK, resp)
	return
}

func (srv *UserService) Update(ctx *gin.Context) {
	panic("implement me")
}

func (srv *UserService) Delete(ctx *gin.Context) {
	panic("implement me")
}
