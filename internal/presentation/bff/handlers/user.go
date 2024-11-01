package handlers

import "github.com/gin-gonic/gin"

// Create       godoc
// @Summary     Creating a new user
// @Tags        User
// @Description Creating a new user
// @Produce     json
// @Param       input body   dto.CreateUserRequest true "User"
// @Success     200          {object} dto.UserBriefResponse
// @Failure     400,404,422  {object} lib.HTTPError
// @Failure     500          {object} lib.HTTPError
// @Failure     default      {object} lib.HTTPError
// @Router      /v1/user [post]
func (h *Handler) Create(ctx *gin.Context) { h.service.Create(ctx) }

// Get           godoc
// @Summary      Get user by ID
// @Tags         User
// @Description  Get user by ID
// @Produce      json
// @Param        user_id path string true "User ID"
// @Success      200                    {object} dto.UserResponse
// @Failure      400,404,422            {object} lib.HTTPError
// @Failure      500                    {object} lib.HTTPError
// @Failure      default                {object} lib.HTTPError
// @Router       /v1/user/{user_id} [get]
func (h *Handler) Get(ctx *gin.Context) { h.service.Get(ctx) }

// Update       godoc
// @Summary     Updating user data
// @Tags        User
// @Description Updating user data
// @Produce     json
// @Param       user_id path string true "User ID"
// @Param       input body             dto.UpdateUserRequest true "User"
// @Success     200                    {object} dto.UserResponse
// @Failure     400,404,422            {object} lib.HTTPError
// @Failure     500                    {object} lib.HTTPError
// @Failure     default                {object} lib.HTTPError
// @Router      /v1/user/{user_id} [put]
func (h *Handler) Update(ctx *gin.Context) { h.service.Update(ctx) }
