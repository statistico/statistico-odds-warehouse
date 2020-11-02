package bootstrap

import (
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

type Container struct {
	Config             *Config
	DatabaseConnection *sql.DB
	Logger             *logrus.Logger
}

func BuildContainer(config *Config) *Container {
	c := Container{
		Config: config,
	}

	c.DatabaseConnection = databaseConnection(config)
	c.Logger = logger()

	return &c
}

func databaseConnection(config *Config) *sql.DB {
	db := config.Database

	dsn := "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable"

	psqlInfo := fmt.Sprintf(dsn, db.Host, db.Port, db.User, db.Password, db.Name)

	conn, err := sql.Open(db.Driver, psqlInfo)

	if err != nil {
		panic(err)
	}

	conn.SetMaxOpenConns(50)
	conn.SetMaxIdleConns(25)

	return conn
}

func logger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)
	return logger
}
