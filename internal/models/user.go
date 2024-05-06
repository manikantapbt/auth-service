package models

import (
	v1 "auth-service/internal/gen/auth/v1"
)

type User struct {
	Id          int32
	Name        string
	UserName    string
	Email       string
	CreatedAt   string
	Verified    bool
	CountryCode int32
	PhoneNumber string
}

func ToUser(request *v1.SignupWithPhoneNumberRequest) *User {
	user := request.User
	return &User{
		Name:        user.Name,
		Email:       user.Email,
		UserName:    user.UserName,
		CountryCode: user.CountryCode,
		PhoneNumber: user.PhoneNumber,
		Verified:    false,
		Id:          0,
	}
}

func ToProto(user *User) *v1.User {
	return &v1.User{
		Id:          user.Id,
		Name:        user.Name,
		UserName:    user.UserName,
		Email:       user.Email,
		IsVerified:  user.Verified,
		CountryCode: user.CountryCode,
		PhoneNumber: user.PhoneNumber,
	}
}
