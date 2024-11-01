package dto

import (
	"github.com/google/uuid"
	"time"
)

type UserBriefRequest struct {
	UserID uuid.UUID `json:"user_id" binding:"required" example:"3a282222-2100-4ff7-9849-71b853da0926"`
}

type CreateUserRequest struct {
	UserName        string `json:"username" binding:"required" example:"JD"`
	FullName        string `json:"full_name" binding:"required" example:"John Dorian"`
	Email           string `json:"email" binding:"required" example:"j.dorian@example.com"`
	Password        string `json:"password" binding:"required" example:"scrubs"`
	ConfirmPassword string `json:"confirm_password" binding:"required" example:"scrubs"`
}

type LoginRequest struct {
	UserName string `json:"username" binding:"required" example:"JD"`
	Password string `json:"password" binding:"required" example:"scrubs"`
}

type UpdateUserRequest struct {
	UserID       uuid.UUID         `json:"user_id" example:"3a282222-2100-4ff7-9849-71b853da0926"`
	UserName     string            `json:"username" example:"JD"`
	FullName     string            `json:"full_name" example:"John Dorian"`
	Email        string            `json:"email" example:"j.dorian@example.com"`
	CustomFields map[string]string `json:"custom_fields" example:"key:value"`
}

type UserBriefResponse struct {
	UserID uuid.UUID `json:"user_id" example:"3a282222-2100-4ff7-9849-71b853da0926"`
}

type UserResponse struct {
	UserID       uuid.UUID         `json:"user_id" example:"3a282222-2100-4ff7-9849-71b853da0926"`
	UserName     string            `json:"username" example:"JD"`
	FullName     string            `json:"full_name" example:"John Dorian"`
	Email        string            `json:"email" example:"j.dorian@example.com"`
	Active       bool              `json:"active" example:"true"`
	Hidden       bool              `json:"hidden" example:"false"`
	LastLogin    time.Time         `json:"last_login" example:"2020-12-31T23:59:59Z"`
	CustomFields map[string]string `json:"custom_fields" example:"key:value"`
	CreatedAt    time.Time         `json:"created_at" example:"2020-12-31T23:59:59Z"`
	UpdatedAt    time.Time         `json:"updated_at" example:"2020-12-31T23:59:59Z"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token" example:"access_token"`
}
