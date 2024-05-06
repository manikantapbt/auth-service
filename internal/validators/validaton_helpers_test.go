package validators

import (
	"testing"
)

func TestValidateEmail(t *testing.T) {
	// Test valid email
	validEmail := "test@example.com"
	if err := validateEmail(validEmail); err != nil {
		t.Errorf("validateEmail(%s) returned error: %v", validEmail, err)
	}

	// Test invalid email
	invalidEmail := "invalid_email.com"
	if err := validateEmail(invalidEmail); err == nil {
		t.Errorf("validateEmail(%s) expected error, but got nil", invalidEmail)
	}
}

func TestValidateName(t *testing.T) {
	// Test valid name
	validName := "John Doe"
	if err := validateName(validName); err != nil {
		t.Errorf("validateName(%s) returned error: %v", validName, err)
	}

	// Test empty name
	emptyName := ""
	if err := validateName(emptyName); err == nil {
		t.Errorf("validateName(%s) expected error, but got nil", emptyName)
	}
}

func TestValidateUserName(t *testing.T) {
	// Test valid user name
	validUserName := "john_doe"
	if err := validateUserName(validUserName); err != nil {
		t.Errorf("validateUserName(%s) returned error: %v", validUserName, err)
	}

	// Test empty user name
	emptyUserName := ""
	if err := validateUserName(emptyUserName); err == nil {
		t.Errorf("validateUserName(%s) expected error, but got nil", emptyUserName)
	}
}

func TestValidatePhoneNumber(t *testing.T) {
	// Test valid phone number
	validPhoneNumber := "+911234567890"
	if err := validatePhoneNumber(validPhoneNumber); err != nil {
		t.Errorf("validatePhoneNumber(%s) returned error: %v", validPhoneNumber, err)
	}

	//Test invalid phone number
	invalidPhoneNumber := "12567890"
	if err := validatePhoneNumber(invalidPhoneNumber); err == nil {
		t.Errorf("validatePhoneNumber(%s) expected error, but got nil", invalidPhoneNumber)
	}
}

func TestValidateCountryCodes(t *testing.T) {
	// Test valid country code
	validCountryCode := int32(91)
	if err := validateCountryCodes(validCountryCode); err != nil {
		t.Errorf("validateCountryCodes(%d) returned error: %v", validCountryCode, err)
	}

	// Test invalid country code
	invalidCountryCode := int32(99)
	if err := validateCountryCodes(invalidCountryCode); err == nil {
		t.Errorf("validateCountryCodes(%d) expected error, but got nil", invalidCountryCode)
	}
}

func TestValidateOtp(t *testing.T) {
	// Test valid OTP
	validOTP := int32(123456)
	if err := validateOtp(validOTP); err != nil {
		t.Errorf("validateOtp(%d) returned error: %v", validOTP, err)
	}

	// Test invalid OTP
	invalidOTP := int32(12345)
	if err := validateOtp(invalidOTP); err == nil {
		t.Errorf("validateOtp(%d) expected error, but got nil", invalidOTP)
	}
}
