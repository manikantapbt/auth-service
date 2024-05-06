package validators

import (
	"auth-service/internal/gen/auth/v1"
	"testing"
)

func TestValidateSignupWithPhoneNumberRequest(t *testing.T) {
	validRequest := &v1.SignupWithPhoneNumberRequest{
		User: &v1.User{
			Name:        "John Doe",
			UserName:    "johndoe",
			Email:       "john@example.com",
			CountryCode: 91,
			PhoneNumber: "1234567890",
		},
	}

	invalidRequest := &v1.SignupWithPhoneNumberRequest{
		User: &v1.User{
			Name:        "", // Empty name
			UserName:    "johndoe",
			Email:       "john@example.com",
			CountryCode: 91,
			PhoneNumber: "+911234567890",
		},
	}

	validator := NewValidator()

	// Test valid request
	if err := validator.ValidateSignupWithPhoneNumberRequest(validRequest); err != nil {
		t.Errorf("ValidateSignupWithPhoneNumberRequest returned error for valid request: %v", err)
	}

	// Test invalid request
	if err := validator.ValidateSignupWithPhoneNumberRequest(invalidRequest); err == nil {
		t.Errorf("ValidateSignupWithPhoneNumberRequest expected error for invalid request, but got nil")
	}
}

func TestValidateLoginWithPhoneNumberRequest(t *testing.T) {
	validRequest := &v1.LoginWithPhoneNumberRequest{
		PhoneNumber: "+911234567890",
		CountryCode: 91,
	}

	invalidRequest := &v1.LoginWithPhoneNumberRequest{
		PhoneNumber: "", // Empty phone number
		CountryCode: 91,
	}

	validator := NewValidator()

	// Test valid request
	if err := validator.ValidateLoginWithPhoneNumberRequest(validRequest); err != nil {
		t.Errorf("ValidateLoginWithPhoneNumberRequest returned error for valid request: %v", err)
	}

	// Test invalid request
	if err := validator.ValidateLoginWithPhoneNumberRequest(invalidRequest); err == nil {
		t.Errorf("ValidateLoginWithPhoneNumberRequest expected error for invalid request, but got nil")
	}
}

func TestValidateVerifyPhoneNumberRequest(t *testing.T) {
	validRequest := &v1.VerifyPhoneNumberRequest{
		PhoneNumber: "+911234567890",
		Otp:         123456,
		CountryCode: 91,
	}

	invalidRequest := &v1.VerifyPhoneNumberRequest{
		PhoneNumber: "", // Empty phone number
		Otp:         123456,
		CountryCode: 91,
	}

	validator := NewValidator()

	// Test valid request
	if err := validator.ValidateVerifyPhoneNumberRequest(validRequest); err != nil {
		t.Errorf("ValidateVerifyPhoneNumberRequest returned error for valid request: %v", err)
	}

	// Test invalid request
	if err := validator.ValidateVerifyPhoneNumberRequest(invalidRequest); err == nil {
		t.Errorf("ValidateVerifyPhoneNumberRequest expected error for invalid request, but got nil")
	}
}

func TestValidatePhoneNumberLogin(t *testing.T) {
	validRequest := &v1.ValidatePhoneNumberLoginRequest{
		PhoneNumber: "+911234567890",
		Otp:         123456,
		CountryCode: 91,
	}

	invalidRequest := &v1.ValidatePhoneNumberLoginRequest{
		PhoneNumber: "", // Empty phone number
		Otp:         123456,
		CountryCode: 91,
	}

	validator := NewValidator()

	// Test valid request
	if err := validator.ValidatePhoneNumberLogin(validRequest); err != nil {
		t.Errorf("ValidatePhoneNumberLogin returned error for valid request: %v", err)
	}

	// Test invalid request
	if err := validator.ValidatePhoneNumberLogin(invalidRequest); err == nil {
		t.Errorf("ValidatePhoneNumberLogin expected error for invalid request, but got nil")
	}
}

func TestValidateGetProfileByMobileNumberRequest(t *testing.T) {
	validRequest := &v1.GetProfileByPhoneNumberRequest{
		PhoneNumber: "+911234567890",
		CountryCode: 91,
	}

	invalidRequest := &v1.GetProfileByPhoneNumberRequest{
		PhoneNumber: "", // Empty phone number
		CountryCode: 91,
	}

	validator := NewValidator()

	// Test valid request
	if err := validator.ValidateGetProfileByMobileNumberRequest(validRequest); err != nil {
		t.Errorf("ValidateGetProfileByMobileNumberRequest returned error for valid request: %v", err)
	}

	// Test invalid request
	if err := validator.ValidateGetProfileByMobileNumberRequest(invalidRequest); err == nil {
		t.Errorf("ValidateGetProfileByMobileNumberRequest expected error for invalid request, but got nil")
	}
}
