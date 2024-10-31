package handlers

import "github.com/gin-gonic/gin"

// Register     godoc
// @Summary     Registering a new user
// @Tags        Auth
// @Description Registering a new user
// @Produce     json
// @Param       input body            dto.RegisterUserRequest true "User"
// @Success     200                   {object} dto.RegisterUserResponse
// @Failure     400,404,422           {object} common.HTTPError
// @Failure     500                   {object} common.HTTPError
// @Failure     default               {object} common.HTTPError
// @Router      /v1/auth/register [post]
func (h *Handler) Register(ctx *gin.Context) { h.service.Register(ctx) }

// Login        godoc
// @Summary     User login to the system
// @Tags        Auth
// @Description User login to the system
// @Produce     json
// @Param       input body         dto.LoginRequest true "Login"
// @Success     200                {object} dto.LoginResponse
// @Failure     400,404,422        {object} common.HTTPError
// @Failure     500                {object} common.HTTPError
// @Failure     default            {object} common.HTTPError
// @Router      /v1/auth/login [post]
func (h *Handler) Login(ctx *gin.Context) { h.service.Login(ctx) }
