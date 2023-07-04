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

type Database struct {
	connectionOptions *ConnectionOptions
	logger            *zap.SugaredLogger
	Client            *gorm.DB
}

func (db *Database) Connect() error {
	host := db.connectionOptions.Host
	port := db.connectionOptions.Port
	name := db.connectionOptions.Name
	username := db.connectionOptions.Username
	password := db.connectionOptions.Password
	ssl := db.connectionOptions.Ssl

	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s", host, port, name, username, password, ssl)
	client, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		return err
	}

	db.Client = client

	return nil
}

func (db *Database) Disconnect() error {
	native, err := db.Client.DB()

	if err != nil {
		return err
	}

	native.Close()

	return nil
}

func New(options *ConnectionOptions, logger *zap.SugaredLogger) *Database {
	return &Database{
		logger:            logger,
		connectionOptions: options,
	}
}
