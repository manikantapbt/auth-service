package service

import (
	"auth-service/internal/gateway"
	auth "auth-service/internal/gen/auth/v1"
	otp "auth-service/internal/gen/otp/v1"
	"auth-service/internal/models"
	"auth-service/internal/repository"
	"auth-service/internal/validators"
	"errors"
	"fmt"
)

type UserEvents string

const (
	SIGN_IN_REQUEST_OTP      UserEvents = "SIGN_REQUEST_OTP"
	PHONE_VERIFIED           UserEvents = "PHONE_VERIFIED"
	INCORRECT_OTP            UserEvents = "INCORRECT_OTP"
	UNVERIFIED_LOGIN_ATTEMPT UserEvents = "UNVERIFIED_LOGIN_ATTEMPT"
	LOGIN_REQUEST            UserEvents = "LOGIN_REQUEST"
	LOGIN_SUCCESSFUL         UserEvents = "LOGIN"
	LOGOUT                   UserEvents = "LOGOUT"
)

type IAuthService interface {
	HandleSignUp(*auth.SignupWithPhoneNumberRequest) (*auth.User, error)
	GetUserProfile(*auth.GetProfileRequest) (*auth.User, error)
	GetUserProfileByPhone(*auth.GetProfileByPhoneNumberRequest) (*auth.User, error)
	VerifyOtp(request *auth.VerifyPhoneNumberRequest) error
	LoginWithPhoneNumber(request *auth.LoginWithPhoneNumberRequest) error
	ValidatePhoneNumberLogin(request *auth.ValidatePhoneNumberLoginRequest) error
}

type authService struct {
	repository.IUserRepository
	validators.IRequestValidator
	publisher gateway.IMessagePublisher
	IGenerator
	repository.IEventRepository
}

func (a authService) HandleSignUp(request *auth.SignupWithPhoneNumberRequest) (*auth.User, error) {
	err := a.ValidateSignupWithPhoneNumberRequest(request)
	if err != nil {
		return nil, err
	}
	user := models.ToUser(request)
	savedUser, err := a.SaveUser(user)
	if err != nil {
		return nil, err
	}
	err = a.publishMessageForOtp(savedUser)
	if err != nil {
		return nil, err
	}
	a.InsertEvent(string(SIGN_IN_REQUEST_OTP), savedUser.PhoneNumber)
	return models.ToProto(savedUser), nil
}

func (a authService) GetUserProfile(request *auth.GetProfileRequest) (*auth.User, error) {
	user, err := a.GetUser(request.UserId)
	if err != nil {
		return nil, err
	}
	return models.ToProto(user), nil
}

func (a authService) GetUserProfileByPhone(request *auth.GetProfileByPhoneNumberRequest) (*auth.User, error) {
	err := a.ValidateGetProfileByMobileNumberRequest(request)
	if err != nil {
		return nil, err
	}
	user, err := a.GetUserByPhoneNumberAndCountry(request.CountryCode, request.PhoneNumber)
	if err != nil {
		return nil, err
	}
	return models.ToProto(user), nil
}

func (a authService) VerifyOtp(request *auth.VerifyPhoneNumberRequest) error {
	err := a.ValidateVerifyPhoneNumberRequest(request)
	if err != nil {
		return err
	}
	user, err := a.GetUserByPhoneNumberAndCountry(request.CountryCode, request.PhoneNumber)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New(fmt.Sprintf("No user registered with %s", request.PhoneNumber))
	}
	generatedOtp, err := a.Generate(request.PhoneNumber)
	if err != nil {
		return errors.New("unable to verify the OTP, Please try again after some time")
	}
	if generatedOtp != request.Otp {
		a.InsertEvent(string(INCORRECT_OTP), user.PhoneNumber)
		return errors.New("invalid OTP")
	}
	err = a.UpdateVerifiedTrueById(user.Id)
	if err != nil {
		return err
	}
	a.InsertEvent(string(PHONE_VERIFIED), user.PhoneNumber)
	return nil
}

func (a authService) LoginWithPhoneNumber(request *auth.LoginWithPhoneNumberRequest) error {
	err := a.ValidateLoginWithPhoneNumberRequest(request)
	if err != nil {
		return err
	}
	user, err := a.GetUserByPhoneNumberAndCountry(request.CountryCode, request.PhoneNumber)
	if err != nil {
		return err
	}
	if !user.Verified {
		a.InsertEvent(string(UNVERIFIED_LOGIN_ATTEMPT), user.PhoneNumber)
		return fmt.Errorf("verify phone number to login")
	}
	err = a.publishMessageForOtp(user)
	if err != nil {
		return err
	}
	a.InsertEvent(string(LOGIN_REQUEST), user.PhoneNumber)
	return nil
}

func (a authService) ValidatePhoneNumberLogin(request *auth.ValidatePhoneNumberLoginRequest) error {
	err := a.IRequestValidator.ValidatePhoneNumberLogin(request)
	if err != nil {
		return err
	}
	user, err := a.GetUserByPhoneNumberAndCountry(request.CountryCode, request.PhoneNumber)
	if err != nil {
		return err
	}
	generatedOtp, err := a.Generate(request.PhoneNumber)
	if err != nil {
		return errors.New("unable to verify the OTP, Please try again after some time")
	}
	if generatedOtp != request.Otp {
		a.InsertEvent(string(INCORRECT_OTP), user.PhoneNumber)
		return errors.New("invalid OTP")
	}
	a.InsertEvent(string(LOGIN_SUCCESSFUL), user.PhoneNumber)
	return nil
}

func (a authService) publishMessageForOtp(user *models.User) error {
	request := &otp.GenerateOTPRequest{
		CountryCode: user.CountryCode,
		PhoneNumber: user.PhoneNumber,
	}
	return a.publisher.Publish(request)
}

func NewAuthService(userRepository repository.IUserRepository, validator validators.IRequestValidator, publisher gateway.IMessagePublisher, generator IGenerator, eventRepository repository.IEventRepository) IAuthService {
	return &authService{IUserRepository: userRepository, IRequestValidator: validator, publisher: publisher, IGenerator: generator, IEventRepository: eventRepository}
}
