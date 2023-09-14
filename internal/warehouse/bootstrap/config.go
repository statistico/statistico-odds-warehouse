package bootstrap

import "os"

type Config struct {
	Database
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

type Sentry struct {
	DSN string
}

func BuildConfig(ssm bool) *Config {
	config := Config{}

	if ssm {
		config.Database = Database{
			Driver:   getSsmParameter("statistico-odds-warehouse-DB_DRIVER"),
			Host:     getSsmParameter("statistico-odds-warehouse-DB_HOST"),
			Port:     getSsmParameter("statistico-odds-warehouse-DB_PORT"),
			User:     getSsmParameter("statistico-odds-warehouse-DB_USER"),
			Password: getSsmParameter("statistico-odds-warehouse-DB_PASSWORD"),
			Name:     getSsmParameter("statistico-odds-warehouse-DB_NAME"),
		}
	} else {
		config.Database = Database{
			Driver:   os.Getenv("DB_DRIVER"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
		}
	}

	if ssm {
		config.Sentry = Sentry{DSN: getSsmParameter("statistico-odds-warehouse-SENTRY_DSN")}
	} else {
		config.Sentry = Sentry{DSN: os.Getenv("SENTRY_DSN")}
	}

	return &config
}
