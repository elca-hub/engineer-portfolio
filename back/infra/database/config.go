package database

import (
	"os"
	"time"
)

type MysqlConfig struct {
	host     string
	database string
	port     string
	user     string
	password string

	ctxTimeout time.Duration
}

type RedisConfig struct {
	host     string
	database string
	port     string
	password string
	poolSize int
}

func NewMySQLConfig() *MysqlConfig {
	return &MysqlConfig{
		host:     os.Getenv("MYSQL_HOST"),
		database: os.Getenv("MYSQL_DATABASE"),
		port:     os.Getenv("MYSQL_PORT"),
		user:     os.Getenv("MYSQL_USER"),
		password: os.Getenv("MYSQL_PASSWORD"),
	}
}

func NewMyNoSQLConfig() *RedisConfig {
	return &RedisConfig{
		host:     os.Getenv("REDIS_HOST"),
		port:     os.Getenv("REDIS_PORT"),
		password: os.Getenv("REDIS_PASSWORD"),
		poolSize: 100,
	}
}
