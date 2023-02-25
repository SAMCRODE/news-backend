package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DATABASEURI string
	JWTKEY      string
	MODE        string
}

var config Config

func Init() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading .env file")
		// log.Fatalf("Error loading .env file")
	}

	// log.Println(os.Getenv("DATABASEURI"))
	// log.Println(config)
	config.DATABASEURI = os.Getenv("DATABASEURI")
	config.JWTKEY = os.Getenv("JWTKEY")
	config.MODE = os.Getenv("MODE")
}

func GetConfig() Config {
	return config
}
