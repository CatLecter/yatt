package schemes

import (
	"time"
)

type UserBriefRequest struct {
	UserID string `json:"user_id" binding:"required"`
}

type UserRequest struct {
	UserName        string `json:"username" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}

type UserBriefResponse struct {
	UserID string `json:"user_id"`
}

type UserResponse struct {
	UserID        string    `json:"user_id"`
	UserName      string    `json:"username"`
	Email         string    `json:"email"`
	LastLoginDate time.Time `json:"last_login_date"`
	IsActive      bool      `json:"is_active"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
