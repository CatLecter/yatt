package domain

import (
	"google.golang.org/grpc/codes"
	"time"
	"yatt/internal/lib"
)

func New() *UserModel {
	return &UserModel{}
}

func Register(username, fullName, email, password, confirmPassword string) (*UserModel, *lib.GRPCError) {
	user := New()
	if err := user.SetUserName(username); err != nil {
		return nil, err
	}
	if err := user.setFullName(fullName); err != nil {
		return nil, err
	}
	if err := user.setEmail(email); err != nil {
		return nil, err
	}
	if err := user.setPassword(password, confirmPassword); err != nil {
		return nil, err
	}
	user.Active = true
	user.Hidden = false
	user.LastLogin = time.Time{}
	return user, nil
}

func (u *UserModel) SetID(id string) *lib.GRPCError {
	u.Mu.Lock()
	defer u.Mu.Unlock()
	if len(id) == 0 {
		return lib.NewGRPCError(codes.InvalidArgument, "empty user_id")
	}
	u.ID = id
	return nil
}

func (u *UserModel) SetUserName(username string) *lib.GRPCError {
	u.Mu.Lock()
	defer u.Mu.Unlock()
	if len(username) == 0 {
		return lib.NewGRPCError(codes.InvalidArgument, "invalid username")
	}
	u.UserName = username
	return nil
}

func (u *UserModel) setFullName(fullName string) *lib.GRPCError {
	if len(fullName) == 0 {
		return lib.NewGRPCError(codes.InvalidArgument, "invalid full_name")
	}
	u.FullName = fullName
	return nil
}

func (u *UserModel) setEmail(email string) *lib.GRPCError {
	if len(email) == 0 {
		return lib.NewGRPCError(codes.InvalidArgument, "invalid email")
	}
	u.Email = email
	return nil
}

func (u *UserModel) setPassword(password, confirmPassword string) *lib.GRPCError {
	if len(password) == 0 {
		return lib.NewGRPCError(codes.InvalidArgument, "invalid password")
	}
	if password != confirmPassword {
		return lib.NewGRPCError(codes.InvalidArgument, "password and confirm password do not match")
	}
	u.Password = lib.HashPassword(password)
	return nil
}

func (u *UserModel) SelfUpdate(username, fullName, email string, customFields map[string]string) {
	u.Mu.Lock()
	defer u.Mu.Unlock()
	if len(username) != 0 {
		u.UserName = username
	}
	if len(fullName) != 0 {
		u.FullName = fullName
	}
	if len(email) != 0 {
		u.Email = email
	}
	if len(customFields) != 0 {
		for key, value := range customFields {
			u.CustomFields[key] = value
		}
	}
}

func (u *UserModel) CreateAccessToken(password string) (string, *lib.GRPCError) {
	if len(password) == 0 {
		return "", lib.NewGRPCError(codes.InvalidArgument, "invalid password")
	}
	if u.Password != lib.HashPassword(password) {
		return "", lib.NewGRPCError(codes.Unauthenticated, "incorrect password")
	}
	accessToken, err := lib.NewAccessToken(u.ID)
	if err != nil {
		return "", lib.NewGRPCError(codes.Unavailable, err.Error())
	}
	return accessToken, nil
}
