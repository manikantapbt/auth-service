package validators

import (
	"errors"
	"fmt"
	"regexp"
)

type CountryCode int32

const (
	India CountryCode = 91
)

var AllowedCountryCodes []CountryCode

func init() {
	AllowedCountryCodes = append(AllowedCountryCodes, India)
}

func validateEmail(email string) error {
	pattern := `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
	matched, _ := regexp.MatchString(pattern, email)
	if !matched {
		return fmt.Errorf("email %s is not a valid email", email)
	}
	return nil
}

func validateName(userName string) error {
	if userName == "" {
		return errors.New("name is empty")
	}
	return nil
}

func validateUserName(userName string) error {
	if userName == "" {
		return errors.New("user name is empty")
	}
	return nil
}

func validatePhoneNumber(phoneNumber string) error {
	pattern := `^\+?\d{0,3}?\d{10}$`
	matched, _ := regexp.MatchString(pattern, phoneNumber)
	if !matched {
		return fmt.Errorf("phone number %s is not a valid number", phoneNumber)
	}
	return nil
}

func validateCountryCodes(countryCode int32) error {
	for _, code := range AllowedCountryCodes {
		if countryCode == int32(code) {
			return nil
		}
	}
	return fmt.Errorf("country code %d is not yet supported", countryCode)
}

func validateOtp(otp int32) error {
	if otp < 100000 || otp > 999999 {
		return errors.New("OTP must be 6 digits long")
	}
	return nil
}
