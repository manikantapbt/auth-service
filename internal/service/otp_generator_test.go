package service_test

import (
	"auth-service/internal/service"
	"testing"
	"time"
)

func TestOtpGenerator_Generate(t *testing.T) {
	phoneNumber := "+911234567890"
	mockKey := "mock-key"
	mockInterval := time.Second * 30

	otpGen := service.NewOtpGenerator(mockKey, mockInterval)

	t.Run("Generate OTP successfully", func(t *testing.T) {
		otp, err := otpGen.Generate(phoneNumber)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if otp < 100000 || otp > 999999 {
			t.Errorf("Expected generated OTP to be between 100000 and 999999, got %d", otp)
		}
	})
}
