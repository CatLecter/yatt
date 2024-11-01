package handlers

import "github.com/gin-gonic/gin"

// Login        godoc
// @Summary     User login to the system
// @Tags        Auth
// @Description User login to the system
// @Produce     json
// @Param       input body         dto.LoginRequest true "Login"
// @Success     200                {object} dto.LoginResponse
// @Failure     400,404,422        {object} lib.HTTPError
// @Failure     500                {object} lib.HTTPError
// @Failure     default            {object} lib.HTTPError
// @Router      /v1/auth/login [post]
func (h *Handler) Login(ctx *gin.Context) { h.service.Login(ctx) }
