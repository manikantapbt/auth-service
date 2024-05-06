package config

import "time"

type Config struct {
	DatabaseConfig DatabaseConfig
	RabbitMQConfig RabbitMQConfig
	OTPConfig      OTPConfig
}

func Load() Config {
	database := DatabaseConfig{
		ConnectionString: "postgresql://postgres:abcd@localhost:5432/postgres?sslmode=disable",
	}
	mq := RabbitMQConfig{
		ConnectionString: "amqps://<username>:<password>@<host>/<virtual_host>",
	}
	config := OTPConfig{
		SecretKey: "your_secret_key",
		Interval:  10 * time.Minute,
	}
	return Config{DatabaseConfig: database, RabbitMQConfig: mq, OTPConfig: config}
}

type DatabaseConfig struct {
	ConnectionString string
}

type RabbitMQConfig struct {
	ConnectionString string
}

type OTPConfig struct {
	SecretKey string
	Interval  time.Duration
}
