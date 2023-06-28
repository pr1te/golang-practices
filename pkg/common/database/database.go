package database

import (
	"fmt"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ConnectionOptions struct {
	Name     string
	Host     string
	Port     string
	Username string
	Password string
	Ssl      string
}

type database struct {
	connectionOptions *ConnectionOptions
	logger            *zap.SugaredLogger
	client            *gorm.DB
}

func (db *database) Connect() error {
	host := db.connectionOptions.Host
	port := db.connectionOptions.Port
	name := db.connectionOptions.Name
	username := db.connectionOptions.Username
	password := db.connectionOptions.Password
	ssl := db.connectionOptions.Ssl

	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s", host, port, name, username, password, ssl)
	client, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	db.client = client

	return nil
}

func (db *database) Disconnect() error {
	native, err := db.client.DB()

	if err != nil {
		return err
	}

	native.Close()

	return nil
}

func New(options *ConnectionOptions, logger *zap.SugaredLogger) *database {
	return &database{
		logger:            logger,
		connectionOptions: options,
	}
}
