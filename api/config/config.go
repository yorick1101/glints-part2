package config

import (
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

var dbConfig *DBConfig

type DBConfig struct {
	UserName string
	Password string
	Host     string
	Port     int
	Name     string
}

func Init() {
	/*
		user := "admin"
		password := "admin"
		host := "localhost"
		portstr := "27017"
		name := "glints"
	*/

	user := checkAndGet("ENV_DB_USER")
	password := checkAndGet("ENV_DB_PASSWORD")
	host := checkAndGet("ENV_DB_HOST")
	portstr := checkAndGet("ENV_DB_PORT")
	name := checkAndGet("ENV_DB_NAME")

	port, err := strconv.Atoi(portstr)
	if err != nil {
		log.Panic("ENV_DB_PORT is not a number", err)
	}

	dbConfig = &DBConfig{
		UserName: user,
		Password: password,
		Host:     host,
		Port:     port,
		Name:     name,
	}
}

func checkAndGet(env string) string {
	value := os.Getenv(env)
	if len(value) == 0 {
		log.Panic("Environment variable ", env, " not found")
	}
	return value
}

func GetDBConfig() *DBConfig {
	return dbConfig
}
