package config

import "time"

type Config struct {
	Database  Database
	RabbitMQ  RabbitMQ
	OTPConfig OTPConfig
}

func Load() Config {
	database := Database{
		ConnectionString: "postgresql://postgres:abcd@localhost:5432/postgres?sslmode=disable",
	}
	mq := RabbitMQ{
		ConnectionString: "amqps://<username>:<password>@<host>/<virtual_host>",
	}
	config := OTPConfig{
		SecretKey: "your_secret_key",
		Interval:  10 * time.Minute,
	}
	return Config{Database: database, RabbitMQ: mq, OTPConfig: config}
}

type Database struct {
	ConnectionString string
}

type RabbitMQ struct {
	ConnectionString string
}

type OTPConfig struct {
	SecretKey string
	Interval  time.Duration
}