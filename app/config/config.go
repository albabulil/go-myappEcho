package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
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

}
