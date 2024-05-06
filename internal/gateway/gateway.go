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
