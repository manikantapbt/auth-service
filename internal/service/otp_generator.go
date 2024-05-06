package service

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"time"
)

type IGenerator interface {
	Generate(phoneNumber string) (int32, error)
}

func NewOtpGenerator(key string, interval time.Duration) IGenerator {
	return &otpGenerator{
		secretKey: key,
		interval:  interval,
	}
}

type otpGenerator struct {
	secretKey string
	interval  time.Duration
}

func (o otpGenerator) Generate(phoneNumber string) (int32, error) {
	return o.generateOtp(phoneNumber)
}

func (o otpGenerator) generateOtp(phoneNumber string) (int32, error) {
	secretKey := o.secretKey
	OTP, err := o.otpHelper(phoneNumber, secretKey)
	if err != nil {
		fmt.Println("Error generating OTP:", err)
		return 0, err
	}

	fmt.Println("OTP:", OTP)
	return OTP, nil
}

func (o otpGenerator) otpHelper(phoneNumber string, secretKey string) (int32, error) {
	now := time.Now().Unix() / int64(o.interval.Seconds())
	message := fmt.Sprintf("%s:%d", phoneNumber, now)
	hash := hmac.New(sha256.New, []byte(secretKey))
	hash.Write([]byte(message))
	hashValue := hash.Sum(nil)
	reader := bytes.NewReader(hashValue)
	var otp uint32
	err := binary.Read(reader, binary.BigEndian, &otp)
	if err != nil {
		return 0, err
	}
	otp = otp % 1000000 // Ensure OTP is at most 6 digits
	if otp < 100000 {
		diff := 100000 - otp
		// adding a prime number as salt
		otp += diff + 137
	}
	return int32(otp), nil
}
