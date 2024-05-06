package service

import (
	auth "auth-service/internal/gen/auth/v1"
	"auth-service/internal/models"
	"auth-service/mocks"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestHandleSignUpSuccess(t *testing.T) {
	mockUserRepo, mockValidator, mockPublisher, _, mockEventRepo, authService := setupAuthServiceMocks(t)
	user := &auth.User{
		Name:        "John Doe",
		UserName:    "johndoe",
		Email:       "john@example.com",
		PhoneNumber: "1234567890",
		CountryCode: 91,
	}
	request := &auth.SignupWithPhoneNumberRequest{User: user}
	mockValidator.On("ValidateSignupWithPhoneNumberRequest", request).Return(nil)
	mockUserRepo.On("SaveUser", mock.Anything).Return(models.ToUser(request), nil)
	mockPublisher.On("Publish", mock.Anything).Return(nil)
	mockEventRepo.On("InsertEvent", string(SIGN_IN_REQUEST_OTP), user.PhoneNumber).Return(nil)
	user, err := authService.HandleSignUp(request)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, request.User.PhoneNumber, user.PhoneNumber)
	mockValidator.AssertCalled(t, "ValidateSignupWithPhoneNumberRequest", request)
	mockUserRepo.AssertCalled(t, "SaveUser", mock.Anything)
	mockEventRepo.AssertCalled(t, "InsertEvent", string(SIGN_IN_REQUEST_OTP), user.PhoneNumber)
	mockValidator.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
	mockPublisher.AssertExpectations(t)
	mockEventRepo.AssertExpectations(t)
}

func TestHandleSignUp_ValidationFailure(t *testing.T) {
	mockValidator := &mocks.IRequestValidator{}
	mockUserRepo := &mocks.IUserRepository{}
	mockPublisher := &mocks.IMessagePublisher{}
	mockGenerator := &mocks.IGenerator{}
	mockEventRepo := &mocks.IEventRepository{}
	authService := NewAuthService(mockUserRepo, mockValidator, mockPublisher, mockGenerator, mockEventRepo)
	user := &auth.User{
		Name:        "John Doe",
		UserName:    "johndoe",
		Email:       "john@example.com",
		PhoneNumber: "1234567890",
		CountryCode: 91,
	}
	request := &auth.SignupWithPhoneNumberRequest{User: user}
	expectedErr := errors.New("validation error")
	mockValidator.On("ValidateSignupWithPhoneNumberRequest", request).Return(expectedErr)
	user, err := authService.HandleSignUp(request)
	assert.Error(t, err)
	assert.Nil(t, user)
	mockValidator.AssertCalled(t, "ValidateSignupWithPhoneNumberRequest", request)
	mockUserRepo.AssertNotCalled(t, "SaveUser", mock.Anything)
	mockPublisher.AssertNotCalled(t, "Publish", mock.Anything)
	mockEventRepo.AssertNotCalled(t, "InsertEvent", mock.Anything, mock.Anything)
	mockValidator.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
	mockPublisher.AssertExpectations(t)
	mockEventRepo.AssertExpectations(t)
}

func TestHandleSignUp_UserSavingFailure(t *testing.T) {
	mockValidator := &mocks.IRequestValidator{}
	mockUserRepo := &mocks.IUserRepository{}
	mockPublisher := &mocks.IMessagePublisher{}
	mockGenerator := &mocks.IGenerator{}
	mockEventRepo := &mocks.IEventRepository{}

	authService := NewAuthService(mockUserRepo, mockValidator, mockPublisher, mockGenerator, mockEventRepo)

	user := &auth.User{
		Name:        "John Doe",
		UserName:    "johndoe",
		Email:       "john@example.com",
		PhoneNumber: "1234567890",
		CountryCode: 91,
	}
	request := &auth.SignupWithPhoneNumberRequest{User: user}
	mockValidator.On("ValidateSignupWithPhoneNumberRequest", request).Return(nil)
	expectedErr := errors.New("user saving error")
	mockUserRepo.On("SaveUser", mock.Anything).Return(nil, expectedErr)
	user, err := authService.HandleSignUp(request)
	assert.Error(t, err)
	assert.Nil(t, user)
	mockValidator.AssertCalled(t, "ValidateSignupWithPhoneNumberRequest", request)
	mockUserRepo.AssertCalled(t, "SaveUser", mock.Anything)
	mockPublisher.AssertNotCalled(t, "Publish", mock.Anything)
	mockEventRepo.AssertNotCalled(t, "InsertEvent", mock.Anything, mock.Anything)
	mockValidator.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
	mockPublisher.AssertExpectations(t)
	mockEventRepo.AssertExpectations(t)
}
func TestHandleSignUp_PublishMessageFailure(t *testing.T) {
	mockValidator := &mocks.IRequestValidator{}
	mockUserRepo := &mocks.IUserRepository{}
	mockPublisher := &mocks.IMessagePublisher{}
	mockGenerator := &mocks.IGenerator{}
	mockEventRepo := &mocks.IEventRepository{}
	authService := NewAuthService(mockUserRepo, mockValidator, mockPublisher, mockGenerator, mockEventRepo)
	user := &auth.User{
		Name:        "John Doe",
		UserName:    "johndoe",
		Email:       "john@example.com",
		PhoneNumber: "1234567890",
		CountryCode: 91,
	}
	request := &auth.SignupWithPhoneNumberRequest{User: user}
	mockValidator.On("ValidateSignupWithPhoneNumberRequest", request).Return(nil)
	mockUserRepo.On("SaveUser", mock.Anything).Return(models.ToUser(request), nil)
	expectedErr := errors.New("publish message error")
	mockPublisher.On("Publish", mock.Anything).Return(expectedErr)
	user, err := authService.HandleSignUp(request)
	assert.Error(t, err)
	assert.Nil(t, user)
	mockValidator.AssertCalled(t, "ValidateSignupWithPhoneNumberRequest", request)
	mockUserRepo.AssertCalled(t, "SaveUser", mock.Anything)
	mockEventRepo.AssertNotCalled(t, "InsertEvent", mock.Anything, mock.Anything)
	mockValidator.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
	mockPublisher.AssertExpectations(t)
	mockEventRepo.AssertExpectations(t)
}

func TestGetUserProfile_Success(t *testing.T) {
	mockUserRepo := &mocks.IUserRepository{}
	mockValidator := &mocks.IRequestValidator{}
	mockPublisher := &mocks.IMessagePublisher{}
	mockGenerator := &mocks.IGenerator{}
	mockEventRepo := &mocks.IEventRepository{}
	authService := NewAuthService(mockUserRepo, mockValidator, mockPublisher, mockGenerator, mockEventRepo)
	mockUser := &models.User{
		Id:          1,
		Name:        "John Doe",
		UserName:    "johndoe",
		Email:       "john@example.com",
		PhoneNumber: "1234567890",
		CountryCode: 91,
		Verified:    true,
	}
	request := &auth.GetProfileRequest{
		RequestId: "123",
		UserId:    1,
	}
	mockUserRepo.On("GetUser", request.UserId).Return(mockUser, nil)
	user, err := authService.GetUserProfile(request)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	expectedUser := &auth.User{
		Id:          mockUser.Id,
		Name:        mockUser.Name,
		UserName:    mockUser.UserName,
		Email:       mockUser.Email,
		IsVerified:  mockUser.Verified,
		CountryCode: mockUser.CountryCode,
		PhoneNumber: mockUser.PhoneNumber,
	}
	assert.Equal(t, expectedUser, user)
	mockUserRepo.AssertCalled(t, "GetUser", request.UserId)
	mockValidator.AssertExpectations(t)
	mockPublisher.AssertExpectations(t)
	mockGenerator.AssertExpectations(t)
	mockEventRepo.AssertExpectations(t)
}

func TestGetUserProfile_Failure(t *testing.T) {
	mockUserRepo := &mocks.IUserRepository{}
	mockValidator := &mocks.IRequestValidator{}
	mockPublisher := &mocks.IMessagePublisher{}
	mockGenerator := &mocks.IGenerator{}
	mockEventRepo := &mocks.IEventRepository{}
	authService := NewAuthService(mockUserRepo, mockValidator, mockPublisher, mockGenerator, mockEventRepo)
	request := &auth.GetProfileRequest{
		RequestId: "123",
		UserId:    1,
	}
	expectedError := errors.New("failed to get user")
	mockUserRepo.On("GetUser", request.UserId).Return(nil, expectedError)
	user, err := authService.GetUserProfile(request)
	assert.Error(t, err)
	assert.EqualError(t, err, expectedError.Error())
	assert.Nil(t, user)
	mockUserRepo.AssertCalled(t, "GetUser", request.UserId)
	mockValidator.AssertExpectations(t)
	mockPublisher.AssertExpectations(t)
	mockGenerator.AssertExpectations(t)
	mockEventRepo.AssertExpectations(t)
}

func TestGetUserProfileByPhone_Success(t *testing.T) {
	mockUserRepo := &mocks.IUserRepository{}
	mockValidator := &mocks.IRequestValidator{}
	mockPublisher := &mocks.IMessagePublisher{}
	mockGenerator := &mocks.IGenerator{}
	mockEventRepo := &mocks.IEventRepository{}
	authService := NewAuthService(mockUserRepo, mockValidator, mockPublisher, mockGenerator, mockEventRepo)
	request := &auth.GetProfileByPhoneNumberRequest{
		RequestId:   "123",
		CountryCode: 91,
		PhoneNumber: "1234567890",
	}
	mockUser := &models.User{
		Id:          1,
		Name:        "John Doe",
		UserName:    "johndoe",
		Email:       "john@example.com",
		CreatedAt:   "2022-05-01",
		Verified:    true,
		CountryCode: 91,
		PhoneNumber: "1234567890",
	}
	mockValidator.On("ValidateGetProfileByMobileNumberRequest", request).Return(nil)
	mockUserRepo.On("GetUserByPhoneNumberAndCountry", request.CountryCode, request.PhoneNumber).Return(mockUser, nil)
	user, err := authService.GetUserProfileByPhone(request)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, mockUser.Id, user.Id)
	assert.Equal(t, mockUser.Name, user.Name)
	assert.Equal(t, mockUser.UserName, user.UserName)
	assert.Equal(t, mockUser.Email, user.Email)
	assert.Equal(t, mockUser.Verified, user.IsVerified)
	assert.Equal(t, mockUser.CountryCode, user.CountryCode)
	assert.Equal(t, mockUser.PhoneNumber, user.PhoneNumber)
	mockValidator.AssertCalled(t, "ValidateGetProfileByMobileNumberRequest", request)
	mockUserRepo.AssertCalled(t, "GetUserByPhoneNumberAndCountry", request.CountryCode, request.PhoneNumber)
	mockPublisher.AssertExpectations(t)
	mockGenerator.AssertExpectations(t)
	mockEventRepo.AssertExpectations(t)
}

func TestGetUserProfileByPhone_Failure(t *testing.T) {
	mockUserRepo := &mocks.IUserRepository{}
	mockValidator := &mocks.IRequestValidator{}
	mockPublisher := &mocks.IMessagePublisher{}
	mockGenerator := &mocks.IGenerator{}
	mockEventRepo := &mocks.IEventRepository{}
	authService := NewAuthService(mockUserRepo, mockValidator, mockPublisher, mockGenerator, mockEventRepo)
	request := &auth.GetProfileByPhoneNumberRequest{
		RequestId:   "123",
		CountryCode: 91,
		PhoneNumber: "1234567890",
	}
	expectedError := errors.New("failed to get user by phone number and country code")
	mockValidator.On("ValidateGetProfileByMobileNumberRequest", request).Return(nil)
	mockUserRepo.On("GetUserByPhoneNumberAndCountry", request.CountryCode, request.PhoneNumber).Return(nil, expectedError)
	user, err := authService.GetUserProfileByPhone(request)
	assert.Error(t, err)
	assert.EqualError(t, err, expectedError.Error())
	assert.Nil(t, user)
	mockValidator.AssertCalled(t, "ValidateGetProfileByMobileNumberRequest", request)
	mockUserRepo.AssertCalled(t, "GetUserByPhoneNumberAndCountry", request.CountryCode, request.PhoneNumber)
	mockPublisher.AssertExpectations(t)
	mockGenerator.AssertExpectations(t)
	mockEventRepo.AssertExpectations(t)
}

func TestGetUserProfileByPhone_ValidationFailure(t *testing.T) {
	mockUserRepo := &mocks.IUserRepository{}
	mockValidator := &mocks.IRequestValidator{}
	mockPublisher := &mocks.IMessagePublisher{}
	mockGenerator := &mocks.IGenerator{}
	mockEventRepo := &mocks.IEventRepository{}
	authService := NewAuthService(mockUserRepo, mockValidator, mockPublisher, mockGenerator, mockEventRepo)
	request := &auth.GetProfileByPhoneNumberRequest{
		RequestId:   "123",
		CountryCode: 91,
		PhoneNumber: "1234567890",
	}
	expectedError := errors.New("request validation failed")
	mockValidator.On("ValidateGetProfileByMobileNumberRequest", request).Return(expectedError)
	user, err := authService.GetUserProfileByPhone(request)
	assert.Error(t, err)
	assert.EqualError(t, err, expectedError.Error())
	assert.Nil(t, user)
	mockValidator.AssertCalled(t, "ValidateGetProfileByMobileNumberRequest", request)
	mockUserRepo.AssertExpectations(t)
	mockPublisher.AssertExpectations(t)
	mockGenerator.AssertExpectations(t)
	mockEventRepo.AssertExpectations(t)
}

func TestVerifyOtp_Success(t *testing.T) {
	mockUserRepo := &mocks.IUserRepository{}
	mockValidator := &mocks.IRequestValidator{}
	mockPublisher := &mocks.IMessagePublisher{}
	mockGenerator := &mocks.IGenerator{}
	mockEventRepo := &mocks.IEventRepository{}
	authService := NewAuthService(mockUserRepo, mockValidator, mockPublisher, mockGenerator, mockEventRepo)
	request := &auth.VerifyPhoneNumberRequest{
		RequestId:   "123",
		Otp:         1234,
		CountryCode: 91,
		PhoneNumber: "1234567890",
	}
	mockValidator.On("ValidateVerifyPhoneNumberRequest", request).Return(nil)
	mockUser := &models.User{
		Id:          1,
		Name:        "John Doe",
		UserName:    "johndoe",
		Email:       "john@example.com",
		Verified:    false,
		CountryCode: 91,
		PhoneNumber: "1234567890",
	}
	mockUserRepo.On("GetUserByPhoneNumberAndCountry", request.CountryCode, request.PhoneNumber).Return(mockUser, nil)
	mockGenerator.On("Generate", request.PhoneNumber).Return(request.Otp, nil)
	mockUserRepo.On("MarkVerified", mockUser.Id).Return(nil)
	mockEventRepo.On("InsertEvent", string(PHONE_VERIFIED), mockUser.PhoneNumber).Return(nil)
	err := authService.VerifyOtp(request)
	assert.NoError(t, err)
	mockValidator.AssertCalled(t, "ValidateVerifyPhoneNumberRequest", request)
	mockUserRepo.AssertCalled(t, "GetUserByPhoneNumberAndCountry", request.CountryCode, request.PhoneNumber)
	mockGenerator.AssertCalled(t, "Generate", request.PhoneNumber)
	mockUserRepo.AssertCalled(t, "MarkVerified", mockUser.Id)
	mockEventRepo.AssertCalled(t, "InsertEvent", string(PHONE_VERIFIED), mockUser.PhoneNumber)
	mockValidator.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
	mockGenerator.AssertExpectations(t)
	mockEventRepo.AssertExpectations(t)
}

func TestVerifyOtp_ValidationFailure(t *testing.T) {
	mockValidator := &mocks.IRequestValidator{}
	authService := NewAuthService(nil, mockValidator, nil, nil, nil)
	request := &auth.VerifyPhoneNumberRequest{RequestId: "123", Otp: 1234, CountryCode: 91, PhoneNumber: "1234567890"}
	expectedErr := errors.New("validation error")
	mockValidator.On("ValidateVerifyPhoneNumberRequest", request).Return(expectedErr)
	err := authService.VerifyOtp(request)
	assert.EqualError(t, err, expectedErr.Error())
	mockValidator.AssertCalled(t, "ValidateVerifyPhoneNumberRequest", request)
	mockValidator.AssertExpectations(t)
}

func TestVerifyOtp_GetUserFailure(t *testing.T) {
	mockValidator := &mocks.IRequestValidator{}
	mockUserRepo := &mocks.IUserRepository{}
	authService := NewAuthService(mockUserRepo, mockValidator, nil, nil, nil)
	request := &auth.VerifyPhoneNumberRequest{RequestId: "123", Otp: 1234, CountryCode: 91, PhoneNumber: "1234567890"}
	expectedErr := errors.New("user not found")
	mockValidator.On("ValidateVerifyPhoneNumberRequest", request).Return(nil)
	mockUserRepo.On("GetUserByPhoneNumberAndCountry", request.CountryCode, request.PhoneNumber).Return(nil, expectedErr)
	err := authService.VerifyOtp(request)
	assert.EqualError(t, err, expectedErr.Error())
	mockValidator.AssertCalled(t, "ValidateVerifyPhoneNumberRequest", request)
	mockUserRepo.AssertCalled(t, "GetUserByPhoneNumberAndCountry", request.CountryCode, request.PhoneNumber)
	mockValidator.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}

func TestVerifyOtp_GetUserNil(t *testing.T) {
	mockValidator := &mocks.IRequestValidator{}
	mockUserRepo := &mocks.IUserRepository{}
	authService := NewAuthService(mockUserRepo, mockValidator, nil, nil, nil)
	request := &auth.VerifyPhoneNumberRequest{RequestId: "123", Otp: 1234, CountryCode: 91, PhoneNumber: "1234567890"}
	mockValidator.On("ValidateVerifyPhoneNumberRequest", request).Return(nil)
	mockUserRepo.On("GetUserByPhoneNumberAndCountry", request.CountryCode, request.PhoneNumber).Return(nil, nil)
	err := authService.VerifyOtp(request)
	assert.EqualError(t, err, fmt.Sprintf("No user registered with %s", request.PhoneNumber))
	mockValidator.AssertCalled(t, "ValidateVerifyPhoneNumberRequest", request)
	mockUserRepo.AssertCalled(t, "GetUserByPhoneNumberAndCountry", request.CountryCode, request.PhoneNumber)
	mockValidator.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}

func TestVerifyOtp_InvalidOtp(t *testing.T) {
	mockUserRepo := &mocks.IUserRepository{}
	mockValidator := &mocks.IRequestValidator{}
	mockEventRepo := &mocks.IEventRepository{}
	mockGenerator := &mocks.IGenerator{}
	authService := NewAuthService(mockUserRepo, mockValidator, nil, mockGenerator, mockEventRepo)
	request := &auth.VerifyPhoneNumberRequest{
		CountryCode: 91,
		PhoneNumber: "1234567890",
		Otp:         123456,
	}
	user := &models.User{
		UserName:    "",
		Id:          1,
		PhoneNumber: request.PhoneNumber,
		CountryCode: request.CountryCode,
	}
	mockValidator.On("ValidateVerifyPhoneNumberRequest", request).Return(nil)
	mockUserRepo.On("GetUserByPhoneNumberAndCountry", request.CountryCode, request.PhoneNumber).Return(user, nil)
	mockGenerator.On("Generate", request.PhoneNumber).Return(int32(654321), nil) // Correct OTP
	mockEventRepo.On("InsertEvent", string(INCORRECT_OTP), request.PhoneNumber).Return(errors.New("failed to insert event"))
	err := authService.VerifyOtp(request)
	assert.Error(t, err)
	mockValidator.AssertCalled(t, "ValidateVerifyPhoneNumberRequest", request)
	mockUserRepo.AssertCalled(t, "GetUserByPhoneNumberAndCountry", request.CountryCode, request.PhoneNumber)
	mockGenerator.AssertCalled(t, "Generate", request.PhoneNumber)
	mockEventRepo.AssertCalled(t, "InsertEvent", string(INCORRECT_OTP), request.PhoneNumber)
	mockValidator.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
	mockGenerator.AssertExpectations(t)
	mockEventRepo.AssertExpectations(t)
}

func TestVerifyOtp_GenerateFailure(t *testing.T) {
	mockUserRepo := &mocks.IUserRepository{}
	mockValidator := &mocks.IRequestValidator{}
	mockEventRepo := &mocks.IEventRepository{}
	mockGenerator := &mocks.IGenerator{}
	authService := NewAuthService(mockUserRepo, mockValidator, nil, mockGenerator, mockEventRepo)
	request := &auth.VerifyPhoneNumberRequest{
		CountryCode: 91,
		PhoneNumber: "1234567890",
		Otp:         123456,
	}
	user := &models.User{
		UserName:    "",
		Id:          1,
		PhoneNumber: request.PhoneNumber,
		CountryCode: request.CountryCode,
	}
	mockValidator.On("ValidateVerifyPhoneNumberRequest", request).Return(nil)
	mockUserRepo.On("GetUserByPhoneNumberAndCountry", request.CountryCode, request.PhoneNumber).Return(user, nil)
	mockGenerator.On("Generate", request.PhoneNumber).Return(int32(0), errors.New("failed to generate OTP"))
	err := authService.VerifyOtp(request)
	assert.Error(t, err)
	mockValidator.AssertCalled(t, "ValidateVerifyPhoneNumberRequest", request)
	mockUserRepo.AssertCalled(t, "GetUserByPhoneNumberAndCountry", request.CountryCode, request.PhoneNumber)
	mockGenerator.AssertCalled(t, "Generate", request.PhoneNumber)
	mockValidator.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
	mockGenerator.AssertExpectations(t)
}

func TestVerifyOtp_UpdateFailure(t *testing.T) {
	mockUserRepo := &mocks.IUserRepository{}
	mockValidator := &mocks.IRequestValidator{}
	mockEventRepo := &mocks.IEventRepository{}
	mockGenerator := &mocks.IGenerator{}
	authService := NewAuthService(mockUserRepo, mockValidator, nil, mockGenerator, mockEventRepo)
	request := &auth.VerifyPhoneNumberRequest{
		CountryCode: 91,
		PhoneNumber: "1234567890",
		Otp:         123456,
	}
	user := &models.User{
		UserName:    "",
		Id:          1,
		PhoneNumber: request.PhoneNumber,
		CountryCode: request.CountryCode,
	}
	mockValidator.On("ValidateVerifyPhoneNumberRequest", request).Return(nil)
	mockUserRepo.On("GetUserByPhoneNumberAndCountry", request.CountryCode, request.PhoneNumber).Return(user, nil)
	mockGenerator.On("Generate", request.PhoneNumber).Return(int32(123456), nil)
	mockUserRepo.On("MarkVerified", user.Id).Return(errors.New("failed to update user verification"))
	err := authService.VerifyOtp(request)
	assert.Error(t, err)
	mockValidator.AssertCalled(t, "ValidateVerifyPhoneNumberRequest", request)
	mockUserRepo.AssertCalled(t, "GetUserByPhoneNumberAndCountry", request.CountryCode, request.PhoneNumber)
	mockGenerator.AssertCalled(t, "Generate", request.PhoneNumber)
	mockUserRepo.AssertCalled(t, "MarkVerified", user.Id)
	mockValidator.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
	mockGenerator.AssertExpectations(t)
}

func TestLoginWithPhoneNumber_Success(t *testing.T) {
	mockUserRepo, mockValidator, mockPublisher, _, mockEventRepo, authService := setupAuthServiceMocks(t)
	request := &auth.LoginWithPhoneNumberRequest{
		CountryCode: 91,
		PhoneNumber: "1234567890",
	}
	mockUser := &models.User{
		Id:          1,
		Name:        "John Doe",
		UserName:    "johndoe",
		Email:       "john@example.com",
		Verified:    true,
		CountryCode: 91,
		PhoneNumber: "1234567890",
	}
	mockValidator.On("ValidateLoginWithPhoneNumberRequest", request).Return(nil)
	mockUserRepo.On("GetUserByPhoneNumberAndCountry", request.CountryCode, request.PhoneNumber).Return(mockUser, nil)
	mockPublisher.On("Publish", mock.Anything).Return(nil)
	mockEventRepo.On("InsertEvent", string(LOGIN_REQUEST), request.PhoneNumber).Return(nil)

	err := authService.LoginWithPhoneNumber(request)

	assert.NoError(t, err)

	mockValidator.AssertCalled(t, "ValidateLoginWithPhoneNumberRequest", request)
	mockUserRepo.AssertCalled(t, "GetUserByPhoneNumberAndCountry", request.CountryCode, request.PhoneNumber)
	mockPublisher.AssertCalled(t, "Publish", mock.Anything)
	mockEventRepo.AssertCalled(t, "InsertEvent", string(LOGIN_REQUEST), request.PhoneNumber)
}

func TestLoginWithPhoneNumber_ValidationFailure(t *testing.T) {
	mockUserRepo, mockValidator, _, _, _, authService := setupAuthServiceMocks(t)
	request := &auth.LoginWithPhoneNumberRequest{
		CountryCode: 91,
		PhoneNumber: "1234567890",
	}
	expectedErr := errors.New("validation error")
	mockValidator.On("ValidateLoginWithPhoneNumberRequest", request).Return(expectedErr)

	err := authService.LoginWithPhoneNumber(request)

	assert.Error(t, err)
	assert.EqualError(t, err, expectedErr.Error())

	mockValidator.AssertCalled(t, "ValidateLoginWithPhoneNumberRequest", request)
	mockUserRepo.AssertNotCalled(t, "GetUserByPhoneNumberAndCountry", mock.Anything, mock.Anything)
}

func TestLoginWithPhoneNumber_GetUserFailure(t *testing.T) {
	mockUserRepo, mockValidator, _, _, _, authService := setupAuthServiceMocks(t)
	request := &auth.LoginWithPhoneNumberRequest{
		CountryCode: 91,
		PhoneNumber: "1234567890",
	}
	expectedErr := errors.New("failed to get user")
	mockValidator.On("ValidateLoginWithPhoneNumberRequest", request).Return(nil)
	mockUserRepo.On("GetUserByPhoneNumberAndCountry", request.CountryCode, request.PhoneNumber).Return(nil, expectedErr)

	err := authService.LoginWithPhoneNumber(request)

	assert.Error(t, err)
	assert.EqualError(t, err, expectedErr.Error())

	mockValidator.AssertCalled(t, "ValidateLoginWithPhoneNumberRequest", request)
	mockUserRepo.AssertCalled(t, "GetUserByPhoneNumberAndCountry", request.CountryCode, request.PhoneNumber)
}

func TestLoginWithPhoneNumber_UnverifiedUser(t *testing.T) {
	mockUserRepo, mockValidator, mockPublisher, _, mockEventRepo, authService := setupAuthServiceMocks(t)
	request := &auth.LoginWithPhoneNumberRequest{
		CountryCode: 91,
		PhoneNumber: "1234567890",
	}
	mockUser := &models.User{
		Id:          1,
		Name:        "John Doe",
		UserName:    "johndoe",
		Email:       "john@example.com",
		Verified:    false, // User is not verified
		CountryCode: 91,
		PhoneNumber: "1234567890",
	}
	mockValidator.On("ValidateLoginWithPhoneNumberRequest", request).Return(nil)
	mockUserRepo.On("GetUserByPhoneNumberAndCountry", request.CountryCode, request.PhoneNumber).Return(mockUser, nil)
	mockEventRepo.On("InsertEvent", string(UNVERIFIED_LOGIN_ATTEMPT), request.PhoneNumber).Return(nil)
	err := authService.LoginWithPhoneNumber(request)
	assert.Error(t, err)
	assert.EqualError(t, err, "verify phone number to login")
	mockValidator.AssertCalled(t, "ValidateLoginWithPhoneNumberRequest", request)
	mockUserRepo.AssertCalled(t, "GetUserByPhoneNumberAndCountry", request.CountryCode, request.PhoneNumber)
	mockPublisher.AssertNotCalled(t, "Publish", mock.Anything)
	mockEventRepo.AssertCalled(t, "InsertEvent", string(UNVERIFIED_LOGIN_ATTEMPT), request.PhoneNumber)
}

func TestLoginWithPhoneNumber_PublishMessageFailure(t *testing.T) {
	mockUserRepo, mockValidator, mockPublisher, _, _, authService := setupAuthServiceMocks(t)
	request := &auth.LoginWithPhoneNumberRequest{
		CountryCode: 91,
		PhoneNumber: "1234567890",
	}
	mockUser := &models.User{
		Id:          1,
		Name:        "John Doe",
		UserName:    "johndoe",
		Email:       "john@example.com",
		Verified:    true,
		CountryCode: 91,
		PhoneNumber: "1234567890",
	}
	mockValidator.On("ValidateLoginWithPhoneNumberRequest", request).Return(nil)
	mockUserRepo.On("GetUserByPhoneNumberAndCountry", request.CountryCode, request.PhoneNumber).Return(mockUser, nil)
	mockPublisher.On("Publish", mock.Anything).Return(errors.New("failed to publish message"))
	err := authService.LoginWithPhoneNumber(request)
	assert.Error(t, err)
	assert.EqualError(t, err, "failed to publish message")
	mockValidator.AssertCalled(t, "ValidateLoginWithPhoneNumberRequest", request)
	mockUserRepo.AssertCalled(t, "GetUserByPhoneNumberAndCountry", request.CountryCode, request.PhoneNumber)
	mockPublisher.AssertCalled(t, "Publish", mock.Anything)
}

func TestValidatePhoneNumberLogin_Success(t *testing.T) {
	mockUserRepo, mockValidator, _, mockGenerator, mockEventRepo, authService := setupAuthServiceMocks(t)
	request := &auth.ValidatePhoneNumberLoginRequest{
		CountryCode: 91,
		PhoneNumber: "1234567890",
		Otp:         123456,
	}
	mockUser := &models.User{
		Id:          1,
		Name:        "John Doe",
		UserName:    "johndoe",
		Email:       "john@example.com",
		Verified:    true,
		CountryCode: 91,
		PhoneNumber: "1234567890",
	}
	mockValidator.On("ValidatePhoneNumberLogin", request).Return(nil)
	mockUserRepo.On("GetUserByPhoneNumberAndCountry", request.CountryCode, request.PhoneNumber).Return(mockUser, nil)
	mockGenerator.On("Generate", request.PhoneNumber).Return(int32(123456), nil)
	mockEventRepo.On("InsertEvent", string(LOGIN_SUCCESSFUL), request.PhoneNumber).Return(nil)

	err := authService.ValidatePhoneNumberLogin(request)

	assert.NoError(t, err)

	mockValidator.AssertCalled(t, "ValidatePhoneNumberLogin", request)
	mockUserRepo.AssertCalled(t, "GetUserByPhoneNumberAndCountry", request.CountryCode, request.PhoneNumber)
	mockGenerator.AssertCalled(t, "Generate", request.PhoneNumber)
	mockEventRepo.AssertCalled(t, "InsertEvent", string(LOGIN_SUCCESSFUL), request.PhoneNumber)
}

func TestValidatePhoneNumberLogin_GenerateFailure(t *testing.T) {
	mockUserRepo, mockValidator, _, mockGenerator, mockEventRepo, authService := setupAuthServiceMocks(t)
	request := &auth.ValidatePhoneNumberLoginRequest{
		CountryCode: 91,
		PhoneNumber: "1234567890",
		Otp:         123456,
	}
	mockValidator.On("ValidatePhoneNumberLogin", request).Return(nil)
	mockUserRepo.On("GetUserByPhoneNumberAndCountry", request.CountryCode, request.PhoneNumber).Return(nil, nil)
	mockGenerator.On("Generate", request.PhoneNumber).Return(int32(0), errors.New("failed to generate OTP"))

	err := authService.ValidatePhoneNumberLogin(request)

	assert.Error(t, err)
	assert.EqualError(t, err, "unable to verify the OTP, Please try again after some time")

	mockValidator.AssertCalled(t, "ValidatePhoneNumberLogin", request)
	mockUserRepo.AssertCalled(t, "GetUserByPhoneNumberAndCountry", request.CountryCode, request.PhoneNumber)
	mockGenerator.AssertCalled(t, "Generate", request.PhoneNumber)
	mockEventRepo.AssertNotCalled(t, "InsertEvent", mock.Anything, mock.Anything)
}

func TestValidatePhoneNumberLogin_InvalidOTP(t *testing.T) {
	mockUserRepo, mockValidator, _, mockGenerator, mockEventRepo, authService := setupAuthServiceMocks(t)
	request := &auth.ValidatePhoneNumberLoginRequest{
		CountryCode: 91,
		PhoneNumber: "1234567890",
		Otp:         123456,
	}
	mockUser := &models.User{
		Id:          1,
		Name:        "John Doe",
		UserName:    "johndoe",
		Email:       "john@example.com",
		Verified:    true,
		CountryCode: 91,
		PhoneNumber: "1234567890",
	}
	mockValidator.On("ValidatePhoneNumberLogin", request).Return(nil)
	mockUserRepo.On("GetUserByPhoneNumberAndCountry", request.CountryCode, request.PhoneNumber).Return(mockUser, nil)
	mockGenerator.On("Generate", request.PhoneNumber).Return(int32(654321), nil) // Correct OTP
	mockEventRepo.On("InsertEvent", string(INCORRECT_OTP), request.PhoneNumber).Return(nil)
	err := authService.ValidatePhoneNumberLogin(request)
	assert.Error(t, err)
	assert.EqualError(t, err, "invalid OTP")
	mockValidator.AssertCalled(t, "ValidatePhoneNumberLogin", request)
	mockUserRepo.AssertCalled(t, "GetUserByPhoneNumberAndCountry", request.CountryCode, request.PhoneNumber)
	mockGenerator.AssertCalled(t, "Generate", request.PhoneNumber)
	mockEventRepo.AssertCalled(t, "InsertEvent", string(INCORRECT_OTP), request.PhoneNumber)
}

func TestValidatePhoneNumberLogin_ValidationFailure(t *testing.T) {
	mockValidator := &mocks.IRequestValidator{}
	authService := NewAuthService(nil, mockValidator, nil, nil, nil)
	request := &auth.ValidatePhoneNumberLoginRequest{
		RequestId:   "123",
		PhoneNumber: "1234567890",
		CountryCode: 91,
		Otp:         123456,
	}
	expectedErr := errors.New("validation error")
	mockValidator.On("ValidatePhoneNumberLogin", request).Return(expectedErr)
	err := authService.ValidatePhoneNumberLogin(request)
	assert.Error(t, err)
	assert.EqualError(t, err, expectedErr.Error())
	mockValidator.AssertCalled(t, "ValidatePhoneNumberLogin", request)
}

func TestValidatePhoneNumberLogin_GetUserFailure(t *testing.T) {
	// Setup
	mockValidator := &mocks.IRequestValidator{}
	mockUserRepo := &mocks.IUserRepository{}
	authService := NewAuthService(mockUserRepo, mockValidator, nil, nil, nil)
	request := &auth.ValidatePhoneNumberLoginRequest{
		RequestId:   "123",
		PhoneNumber: "1234567890",
		CountryCode: 91,
		Otp:         123456,
	}
	expectedErr := errors.New("failed to get user")
	mockValidator.On("ValidatePhoneNumberLogin", request).Return(nil)
	mockUserRepo.On("GetUserByPhoneNumberAndCountry", request.CountryCode, request.PhoneNumber).Return(nil, expectedErr)
	err := authService.ValidatePhoneNumberLogin(request)
	assert.Error(t, err)
	assert.EqualError(t, err, expectedErr.Error())
	mockValidator.AssertCalled(t, "ValidatePhoneNumberLogin", request)
	mockUserRepo.AssertCalled(t, "GetUserByPhoneNumberAndCountry", request.CountryCode, request.PhoneNumber)
}

func setupAuthServiceMocks(t *testing.T) (*mocks.IUserRepository, *mocks.IRequestValidator, *mocks.IMessagePublisher, *mocks.IGenerator, *mocks.IEventRepository, IAuthService) {
	mockUserRepo := &mocks.IUserRepository{}
	mockValidator := &mocks.IRequestValidator{}
	mockPublisher := &mocks.IMessagePublisher{}
	mockGenerator := &mocks.IGenerator{}
	mockEventRepo := &mocks.IEventRepository{}
	authService := NewAuthService(mockUserRepo, mockValidator, mockPublisher, mockGenerator, mockEventRepo)
	return mockUserRepo, mockValidator, mockPublisher, mockGenerator, mockEventRepo, authService
}
