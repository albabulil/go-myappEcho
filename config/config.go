package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	APP_NAME string
	APP_PORT string
	// jwt key
	JWT_KEY string
	// config database
	DB_HOST    string
	DB_PORT    string
	DB_USER    string
	DB_PASS    string
	DB_NAME    string
	DB_TIMEOUT string
	DB_DEBUG   string
)

func Init() {
	// get env type
	EnvMode := os.Getenv("NODE_ENV")
	if EnvMode == "" {
		EnvMode = "development"
	}

	fmt.Printf("reading environment variable for env type: %s\n", EnvMode)
	err := godotenv.Load("env/" + EnvMode + ".env")
	if err != nil {
		// log with panic, exit service
		log.Panic(err)
	}

	APP_NAME = os.Getenv("APP_NAME")
	APP_PORT = os.Getenv("APP_PORT")

	JWT_KEY = os.Getenv("JWT_KEY")

	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_USER = os.Getenv("DB_USER")
	DB_PASS = os.Getenv("DB_PASS")
	DB_NAME = os.Getenv("DB_NAME")
	DB_TIMEOUT = os.Getenv("DB_TIMEOUT")
	DB_DEBUG = os.Getenv("DB_DEBUG")
}
