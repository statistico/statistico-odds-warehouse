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
		Driver:   getSsmParameter("statistico-odds-warehouse-DB_DRIVER"),
		Host:     getSsmParameter("statistico-odds-warehouse-DB_HOST"),
		Port:     getSsmParameter("statistico-odds-warehouse-DB_PORT"),
		User:     getSsmParameter("statistico-odds-warehouse-DB_USER"),
		Password: getSsmParameter("statistico-odds-warehouse-DB_PASSWORD"),
		Name:     getSsmParameter("statistico-odds-warehouse-DB_NAME"),
	}

	config.AwsConfig = AwsConfig{
		Key:      os.Getenv("AWS_KEY"),
		Secret:   os.Getenv("AWS_SECRET"),
		Region:   os.Getenv("AWS_REGION"),
		QueueUrl: os.Getenv("AWS_QUEUE_URL"),
	}

	config.Sentry = Sentry{DSN: getSsmParameter("statistico-odds-warehouse-SENTRY_DSN")}

	return &config
}
