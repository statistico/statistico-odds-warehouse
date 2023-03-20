package bootstrap

import "os"

type Config struct {
	AwsConfig
	Database
	QueueDriver string
	Sentry
}

type Database struct {
	Driver   string
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type AwsConfig struct {
	Key      string
	Secret   string
	Region   string
	QueueUrl string
}

type Sentry struct {
	DSN string
}

func BuildConfig() *Config {
	config := Config{}

	config.QueueDriver = os.Getenv("QUEUE_DRIVER")

	config.Database = Database{
		Driver:   os.Getenv("DB_DRIVER"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
	}

	config.AwsConfig = AwsConfig{
		Key:      os.Getenv("AWS_KEY"),
		Secret:   os.Getenv("AWS_SECRET"),
		Region:   os.Getenv("AWS_REGION"),
		QueueUrl: os.Getenv("AWS_QUEUE_URL"),
	}

	config.Sentry = Sentry{DSN: os.Getenv("SENTRY_DSN")}

	return &config
}
