package server

import (
	v1 "auth-service/internal/gen/auth/v1"
	"auth-service/internal/service"
	"connectrpc.com/connect"
	"context"
)

type AuthServer struct {
	service service.IAuthService
}

func NewAuthServer(authService service.IAuthService) *AuthServer {
	return &AuthServer{
		service: authService,
	}
}
func (a *AuthServer) SignupWithPhoneNumber(ctx context.Context, req *connect.Request[v1.SignupWithPhoneNumberRequest]) (*connect.Response[v1.SignupWithPhoneNumberResponse], error) {
	response := &v1.SignupWithPhoneNumberResponse{}
	user, err := a.service.HandleSignUp(req.Msg)
	if err != nil {
		response.Error = &v1.Error{
			Message:   err.Error(),
			ErrorCode: 1,
		}
		response.IsSuccess = false
	} else {
		response.IsSuccess = true
		response.UserId = user.Id
	}
	return connect.NewResponse(response), nil
}

func (a *AuthServer) VerifyPhoneNumber(ctx context.Context, request *connect.Request[v1.VerifyPhoneNumberRequest]) (*connect.Response[v1.VerifyPhoneNumberResponse], error) {
	response := &v1.VerifyPhoneNumberResponse{}
	err := a.service.VerifyOtp(request.Msg)
	if err != nil {
		response.Error = &v1.Error{
			Message:   err.Error(),
			ErrorCode: 1,
		}
		response.IsSuccess = false
	} else {
		response.IsSuccess = true
	}
	return connect.NewResponse(response), nil
}

func (a *AuthServer) LoginWithPhoneNumber(ctx context.Context, request *connect.Request[v1.LoginWithPhoneNumberRequest]) (*connect.Response[v1.LoginWithPhoneNumberResponse], error) {
	response := &v1.LoginWithPhoneNumberResponse{}
	err := a.service.LoginWithPhoneNumber(request.Msg)
	if err != nil {
		response.Error = &v1.Error{
			Message:   err.Error(),
			ErrorCode: 1,
		}
		response.IsSuccess = false
	} else {
		response.IsSuccess = true
	}
	return connect.NewResponse(response), nil
}

func (a *AuthServer) ValidatePhoneNumberLogin(ctx context.Context, request *connect.Request[v1.ValidatePhoneNumberLoginRequest]) (*connect.Response[v1.ValidatePhoneNumberLoginResponse], error) {
	response := &v1.ValidatePhoneNumberLoginResponse{}
	err := a.service.ValidatePhoneNumberLogin(request.Msg)
	if err != nil {
		response.Error = &v1.Error{
			Message:   err.Error(),
			ErrorCode: 1,
		}
		response.IsSuccess = false
	} else {
		response.IsSuccess = true
	}
	return connect.NewResponse(response), nil
}

func (a *AuthServer) GetProfile(ctx context.Context, req *connect.Request[v1.GetProfileRequest]) (*connect.Response[v1.GetProfileResponse], error) {
	response := &v1.GetProfileResponse{}
	user, err := a.service.GetUserProfile(req.Msg)
	if err != nil {
		response.Error = &v1.Error{
			Message:   err.Error(),
			ErrorCode: 1,
		}
		response.IsSuccess = false
	} else {
		response.IsSuccess = true
		response.User = user
	}
	return connect.NewResponse(response), nil
}

func (a *AuthServer) GetProfileByPhoneNumber(ctx context.Context, req *connect.Request[v1.GetProfileByPhoneNumberRequest]) (*connect.Response[v1.GetProfileByPhoneNumberResponse], error) {
	response := &v1.GetProfileByPhoneNumberResponse{}
	user, err := a.service.GetUserProfileByPhone(req.Msg)
	if err != nil {
		response.Error = &v1.Error{
			Message:   err.Error(),
			ErrorCode: 1,
		}
		response.IsSuccess = false
	} else {
		response.IsSuccess = true
		response.User = user
	}
	return connect.NewResponse(response), nil
}
