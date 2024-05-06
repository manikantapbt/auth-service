package server

import (
	auth "auth-service/internal/gen/auth/v1"
	_ "auth-service/internal/service"
	"auth-service/mocks"
	"connectrpc.com/connect"
	_ "connectrpc.com/connect"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthServer_HandleSignUp_Success(t *testing.T) {
	mockService := &mocks.IAuthService{}
	authServer := NewAuthServer(mockService)
	User := &auth.User{
		Id:          123,
		Name:        "John Doe",
		UserName:    "johndoe",
		Email:       "john@example.com",
		IsVerified:  false,
		CountryCode: 91,
		PhoneNumber: "1234567890",
	}
	request := &auth.SignupWithPhoneNumberRequest{User: User}
	mockService.On("HandleSignUp", request).Return(User, nil)
	response, err := authServer.SignupWithPhoneNumber(context.Background(), connect.NewRequest(request))
	assert.NoError(t, err)
	assert.True(t, response.Msg.IsSuccess)
}

func TestAuthServer_HandleSignUp_Failure(t *testing.T) {
	mockService := &mocks.IAuthService{}
	authServer := NewAuthServer(mockService)
	request := &auth.SignupWithPhoneNumberRequest{
		User: &auth.User{},
	}
	mockService.On("HandleSignUp", request).Return(nil, errors.New("service call failed"))
	response, _ := authServer.SignupWithPhoneNumber(context.Background(), connect.NewRequest(request))
	assert.False(t, response.Msg.IsSuccess)
}

func TestAuthServer_VerifyPhoneNumber_Success(t *testing.T) {
	mockService := &mocks.IAuthService{}
	authServer := NewAuthServer(mockService)
	request := &auth.VerifyPhoneNumberRequest{
		RequestId:   "123",
		Otp:         123456,
		CountryCode: 1,
		PhoneNumber: "+1234567890",
	}
	mockService.On("VerifyOtp", request).Return(nil)
	response, err := authServer.VerifyPhoneNumber(context.Background(), connect.NewRequest(request))
	assert.NoError(t, err)
	assert.True(t, response.Msg.IsSuccess)
}

func TestAuthServer_VerifyPhoneNumber_Error(t *testing.T) {
	mockService := &mocks.IAuthService{}
	authServer := NewAuthServer(mockService)
	request := &auth.VerifyPhoneNumberRequest{
		RequestId:   "123",
		Otp:         123456,
		CountryCode: 1,
		PhoneNumber: "+1234567890",
	}
	mockService.On("VerifyOtp", request).Return(errors.New("service failed"))
	response, _ := authServer.VerifyPhoneNumber(context.Background(), connect.NewRequest(request))
	assert.False(t, response.Msg.IsSuccess)
}

func TestAuthServer_LoginWithPhoneNumber_Success(t *testing.T) {
	mockService := &mocks.IAuthService{}
	authServer := NewAuthServer(mockService)
	request := &auth.LoginWithPhoneNumberRequest{
		CountryCode: 1,
		PhoneNumber: "+1234567890",
	}
	mockService.On("LoginWithPhoneNumber", request).Return(nil)
	response, err := authServer.LoginWithPhoneNumber(context.Background(), connect.NewRequest(request))
	assert.NoError(t, err)
	assert.True(t, response.Msg.IsSuccess)
}

func TestAuthServer_LoginWithPhoneNumber_Error(t *testing.T) {
	mockService := &mocks.IAuthService{}
	authServer := NewAuthServer(mockService)
	request := &auth.LoginWithPhoneNumberRequest{
		CountryCode: 1,
		PhoneNumber: "+1234567890",
	}
	mockService.On("LoginWithPhoneNumber", request).Return(errors.New("service failed"))
	response, _ := authServer.LoginWithPhoneNumber(context.Background(), connect.NewRequest(request))
	assert.False(t, response.Msg.IsSuccess)
}

func TestAuthServer_ValidatePhoneNumberLogin_Success(t *testing.T) {
	mockService := &mocks.IAuthService{}
	authServer := NewAuthServer(mockService)
	request := &auth.ValidatePhoneNumberLoginRequest{
		CountryCode: 1,
		PhoneNumber: "+1234567890",
		Otp:         123456,
	}
	mockService.On("ValidatePhoneNumberLogin", request).Return(nil)
	response, err := authServer.ValidatePhoneNumberLogin(context.Background(), connect.NewRequest(request))
	assert.NoError(t, err)
	assert.True(t, response.Msg.IsSuccess)
}

func TestAuthServer_ValidatePhoneNumberLogin_Error(t *testing.T) {
	mockService := &mocks.IAuthService{}
	authServer := NewAuthServer(mockService)
	request := &auth.ValidatePhoneNumberLoginRequest{
		CountryCode: 1,
		PhoneNumber: "+1234567890",
		Otp:         123456,
	}
	mockService.On("ValidatePhoneNumberLogin", request).Return(errors.New("service failed"))
	response, _ := authServer.ValidatePhoneNumberLogin(context.Background(), connect.NewRequest(request))
	assert.False(t, response.Msg.IsSuccess)
}

func TestAuthServer_GetProfile_Success(t *testing.T) {
	mockService := &mocks.IAuthService{}
	authServer := NewAuthServer(mockService)
	user := &auth.User{
		Id:          123,
		Name:        "John Doe",
		UserName:    "johndoe",
		Email:       "john@example.com",
		IsVerified:  false,
		CountryCode: 91,
		PhoneNumber: "1234567890",
	}
	request := &auth.GetProfileRequest{}
	mockService.On("GetUserProfile", request).Return(user, nil)
	response, err := authServer.GetProfile(context.Background(), connect.NewRequest(request))
	assert.NoError(t, err)
	assert.True(t, response.Msg.IsSuccess)
	assert.Equal(t, user, response.Msg.User)
}

func TestAuthServer_GetProfile_Error(t *testing.T) {
	mockService := &mocks.IAuthService{}
	authServer := NewAuthServer(mockService)
	request := &auth.GetProfileRequest{}
	mockService.On("GetUserProfile", request).Return(nil, errors.New("service failed"))
	response, _ := authServer.GetProfile(context.Background(), connect.NewRequest(request))
	assert.False(t, response.Msg.IsSuccess)
}

func TestAuthServer_GetProfileByPhoneNumber_Success(t *testing.T) {
	mockService := &mocks.IAuthService{}
	authServer := NewAuthServer(mockService)
	user := &auth.User{
		Id:          123,
		Name:        "John Doe",
		UserName:    "johndoe",
		Email:       "john@example.com",
		IsVerified:  false,
		CountryCode: 91,
		PhoneNumber: "1234567890",
	}
	request := &auth.GetProfileByPhoneNumberRequest{
		CountryCode: 91,
		PhoneNumber: "1234567890",
	}
	mockService.On("GetUserProfileByPhone", request).Return(user, nil)
	response, err := authServer.GetProfileByPhoneNumber(context.Background(), connect.NewRequest(request))
	assert.NoError(t, err)
	assert.True(t, response.Msg.IsSuccess)
	assert.Equal(t, user, response.Msg.User)
}

func TestAuthServer_GetProfileByPhoneNumber_Error(t *testing.T) {
	mockService := &mocks.IAuthService{}
	authServer := NewAuthServer(mockService)
	request := &auth.GetProfileByPhoneNumberRequest{
		CountryCode: 91,
		PhoneNumber: "1234567890",
	}
	mockService.On("GetUserProfileByPhone", request).Return(nil, errors.New("service failed"))
	response, _ := authServer.GetProfileByPhoneNumber(context.Background(), connect.NewRequest(request))
	assert.False(t, response.Msg.IsSuccess)
}
