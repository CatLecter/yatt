package domain

import (
	"github.com/CatLecter/yatt/internal/lib"
	"google.golang.org/grpc/codes"
	"time"
)

func New() *UserModel {
	return &UserModel{}
}

func Register(username, fullName, email, password, confirmPassword string) (*UserModel, *lib.GRPCError) {
	user := New()
	if err := user.SetUserName(username); err != nil {
		return nil, err
	}
	if err := user.SetFullName(fullName); err != nil {
		return nil, err
	}
	if err := user.SetEmail(email); err != nil {
		return nil, err
	}
	if err := user.SetPassword(password, confirmPassword); err != nil {
		return nil, err
	}
	user.Active = true
	user.Hidden = false
	user.LastLogin = time.Time{}
	return user, nil
}

func (u *UserModel) SetID(id string) *lib.GRPCError {
	if len(id) == 0 {
		return lib.NewGRPCError(codes.InvalidArgument, "empty user_id")
	}
	u.ID = id
	return nil
}

func (u *UserModel) SetUserName(username string) *lib.GRPCError {
	if len(username) == 0 {
		return lib.NewGRPCError(codes.InvalidArgument, "invalid username")
	}
	u.UserName = username
	return nil
}

func (u *UserModel) SetFullName(fullName string) *lib.GRPCError {
	if len(fullName) == 0 {
		return lib.NewGRPCError(codes.InvalidArgument, "invalid full_name")
	}
	u.FullName = fullName
	return nil
}

func (u *UserModel) SetEmail(email string) *lib.GRPCError {
	if len(email) == 0 {
		return lib.NewGRPCError(codes.InvalidArgument, "invalid email")
	}
	u.Email = email
	return nil
}

func (u *UserModel) SetPassword(password, confirmPassword string) *lib.GRPCError {
	if len(password) == 0 {
		return lib.NewGRPCError(codes.InvalidArgument, "invalid password")
	}
	if password != confirmPassword {
		return lib.NewGRPCError(codes.InvalidArgument, "password and confirm password do not match")
	}
	u.Password = lib.HashPassword(password)
	return nil
}

func (u *UserModel) UpdateUserName(username string) {
	if len(username) != 0 {
		u.UserName = username
	}
}

func (u *UserModel) UpdateFullName(fullName string) {
	if len(fullName) != 0 {
		u.FullName = fullName
	}
}

func (u *UserModel) UpdateEmail(email string) {
	if len(email) != 0 {
		u.Email = email
	}
}

func (u *UserModel) UpdateCustomFields(customFields map[string]string) {
	if len(customFields) != 0 {
		for key, value := range customFields {
			u.CustomFields[key] = value
		}
	}
}

func (u *UserModel) SelfUpdate(username, fullName, email string, customFields map[string]string) {
	u.UpdateUserName(username)
	u.UpdateFullName(fullName)
	u.UpdateEmail(email)
	u.UpdateCustomFields(customFields)
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
