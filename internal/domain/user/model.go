package domain

import (
	"sync"
	"time"
)

type UserModel struct {
	Mu           sync.RWMutex
	ID           string
	UserName     string
	FullName     string
	Email        string
	Password     string
	Active       bool
	Hidden       bool
	LastLogin    time.Time
	CustomFields map[string]string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
