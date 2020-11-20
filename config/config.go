package config

import "os"

type Database struct {
	Host     string
	User     string
	Password string
	Port     string
	Dbname   string
}

type Redis struct {
	Host string
	Port string
}

type Config struct {
	Database Database
	Redis    Redis
	Port     string
}

/*
NewConfig provide object for config
*/
func NewConfig() *Config {
	return &Config{
		Database: Database{
			Host:     os.Getenv("DB_HOST"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Port:     os.Getenv("DB_PORT"),
			Dbname:   os.Getenv("DB_NAME"),
		},
		Redis: Redis{
			Host: os.Getenv("REDIS_HOST"),
			Port: os.Getenv("REDIS_PORT"),
		},
		Port: os.Getenv("PORT"),
	}
}
