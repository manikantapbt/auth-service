package gateway

import (
	otp "auth-service/internal/gen/otp/v1"
	"github.com/streadway/amqp"
	"google.golang.org/protobuf/proto"
)

type IMessagePublisher interface {
	Publish(request *otp.GenerateOTPRequest) error
}

func NewRabbitMqPublisher(queueName string, channel *amqp.Channel) IMessagePublisher {
	return &rabbitMqPublisher{
		queueName: queueName,
		channel:   channel,
	}
}

type rabbitMqPublisher struct {
	queueName string
	channel   *amqp.Channel
}

func (r rabbitMqPublisher) Publish(request *otp.GenerateOTPRequest) error {
	//generateOtp(request.PhoneNumber)
	marshalledBytes, err := proto.Marshal(request)
	if err != nil {
		return err
	}
	err = r.channel.Publish("", r.queueName, false, false, amqp.Publishing{
		ContentType: "application/octet-stream",
		Body:        marshalledBytes,
	})
	return err
}

type RabbitMQConnectionDetails struct {
	Queue   *amqp.Queue
	Channel *amqp.Channel
}

//func generateOtp(phoneNumber string) {
//	secretKey := "your_secret_key"
//	interval := 10 * time.Minute // 10 minutes interval
//
//	OTP, err := otpHelper(phoneNumber, secretKey, interval)
//	if err != nil {
//		fmt.Println("Error generating OTP:", err)
//		return
//	}
//	fmt.Println("OTP:", OTP)
//}
//
//func otpHelper(userID string, secretKey string, interval time.Duration) (uint32, error) {
//	now := time.Now().Unix() / int64(interval.Seconds())
//	message := fmt.Sprintf("%s:%d", userID, now)
//	hash := hmac.New(sha256.New, []byte(secretKey))
//	hash.Write([]byte(message))
//	hashValue := hash.Sum(nil)
//	reader := bytes.NewReader(hashValue)
//	var otp uint32
//	err := binary.Read(reader, binary.BigEndian, &otp)
//	if err != nil {
//		return 0, err
//	}
//	return otp % 1000000, nil
//}
