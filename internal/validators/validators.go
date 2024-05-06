package validators

import (
	v1 "auth-service/internal/gen/auth/v1"
	"errors"
)

type IRequestValidator interface {
	ValidateSignupWithPhoneNumberRequest(request *v1.SignupWithPhoneNumberRequest) error
	ValidateLoginWithPhoneNumberRequest(request *v1.LoginWithPhoneNumberRequest) error
	ValidateVerifyPhoneNumberRequest(request *v1.VerifyPhoneNumberRequest) error
	ValidatePhoneNumberLogin(request *v1.ValidatePhoneNumberLoginRequest) error
	ValidateGetProfileByMobileNumberRequest(request *v1.GetProfileByPhoneNumberRequest) error
}

func NewValidator() IRequestValidator {
	return &validator{}
}

type validator struct{}

func (v *validator) ValidateSignupWithPhoneNumberRequest(request *v1.SignupWithPhoneNumberRequest) error {
	phoneErr := validatePhoneNumber(request.User.PhoneNumber)
	userErr := validateUserName(request.User.Name)
	userNameErr := validateUserName(request.User.UserName)
	emailErr := validateEmail(request.User.Email)
	countryErr := validateCountryCodes(request.User.CountryCode)
	return errors.Join(phoneErr, userErr, userNameErr, emailErr, countryErr)
}

func (v *validator) ValidateLoginWithPhoneNumberRequest(request *v1.LoginWithPhoneNumberRequest) error {
	phoneErr := validatePhoneNumber(request.PhoneNumber)
	userErr := validateUserName(request.UserId) // TODO: not necessary?
	countryErr := validateCountryCodes(request.CountryCode)
	return errors.Join(phoneErr, userErr, countryErr)
}

func (v *validator) ValidateVerifyPhoneNumberRequest(request *v1.VerifyPhoneNumberRequest) error {
	phoneErr := validatePhoneNumber(request.PhoneNumber)
	otpErr := validateOtp(request.Otp)
	countryErr := validateCountryCodes(request.CountryCode)
	return errors.Join(phoneErr, otpErr, countryErr)
}

func (v *validator) ValidatePhoneNumberLogin(request *v1.ValidatePhoneNumberLoginRequest) error {
	otpErr := validateOtp(request.Otp)
	countryErr := validateCountryCodes(request.CountryCode)
	phError := validatePhoneNumber(request.PhoneNumber)
	return errors.Join(otpErr, countryErr, phError)
}

func (v *validator) ValidateGetProfileByMobileNumberRequest(request *v1.GetProfileByPhoneNumberRequest) error {
	//phoneErr := validatePhoneNumber(request.PhoneNumber)
	// TODO: request id validation
	countryErr := validateCountryCodes(request.CountryCode)
	phoneErr := validatePhoneNumber(request.PhoneNumber)
	return errors.Join(countryErr, phoneErr)
}
