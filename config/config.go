package config

import (
	"fmt"
)

type Config struct {
	Database Database
}

//database configuration mongoDB

type Database struct {
	Host     string
	Port     string
	Username string
	Password string
	Name     string
}

func (db Database) GetURI() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s/%s",
		db.Username,
		db.Password,
		db.Host,
		db.Port,
		db.Name,
	)
}

func GetConfig() Config {
	return Config{
		Database: Database{
			Host:     "localhost",
			Port:     "270",
			Username: "",
			Password: "",
			Name:     "test",
		},
	}
}
