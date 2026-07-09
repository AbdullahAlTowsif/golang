package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configurations Config

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load the env variables:", err)
		os.Exit(1)
	}

	VERSION := os.Getenv("VERSION")
	if VERSION == ""{
		log.Fatal("Version is required")
		os.Exit(1)
	}

	SERVICE_NAME := os.Getenv("SERVICE_NAME")
	if SERVICE_NAME == "" {
		fmt.Println("SERVICE_NAME is required")
		os.Exit(1)
	}

	HTTP_PORT := os.Getenv("HTTP_PORT")
	if HTTP_PORT == "" {
		fmt.Println("HTTP_PORT is required")
		os.Exit(1)
	}

	port, err := strconv.ParseInt(HTTP_PORT, 10, 64)
	if err != nil {
		fmt.Println("Port must be number")
		os.Exit(1)
	}

	configurations = Config{
		Version: VERSION,
		ServiceName: SERVICE_NAME,
		HttpPort: int(port), // typecasting
	}
}

func GetConfig() Config {
	loadConfig()
	return configurations
}