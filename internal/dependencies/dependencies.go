package dependencies

import (
	"auth-service/internal/config"
	"auth-service/internal/gateway"
	"auth-service/internal/repository"
	"auth-service/internal/service"
	"auth-service/internal/validators"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/streadway/amqp"
	"log"
)

type Dependencies struct {
	Db                 *sql.DB
	Service            service.IAuthService
	RabbitMQConnection *amqp.Connection
	Channel            *amqp.Channel
	GateWayService     gateway.IMessagePublisher
}

func Initialize(config config.Config) (*Dependencies, error) {
	db, err := sql.Open("postgres", config.Database.ConnectionString)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	conn, err := amqp.Dial(config.RabbitMQ.ConnectionString)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	queue, err := ch.QueueDeclare(
		"teja", // name
		false,  // durable
		false,  // delete when unused
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
	)
	publisher := gateway.NewRabbitMqPublisher(queue.Name, ch)
	validator := validators.NewValidator()
	newRepository := repository.NewUserRepository(db)
	eventRepository := repository.NewEventRepository(db)
	generator := service.NewOtpGenerator(config.OTPConfig.SecretKey, config.OTPConfig.Interval)
	authService := service.NewAuthService(newRepository, validator, publisher, generator, eventRepository)
	return &Dependencies{
		Db:                 db,
		Service:            authService,
		RabbitMQConnection: conn,
		Channel:            ch,
	}, nil
}

func (d Dependencies) ShutDown() error {
	err := d.Db.Close()
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Println("Database connection closed")
	err = d.Channel.Close()
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Println("RabbitMQ Channel closed")
	err = d.RabbitMQConnection.Close()
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Println("RabbitMQ connection closed")
	return nil
}
