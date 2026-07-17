package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configurations *Config

type DBConfig struct {
	Host          string
	Port          int
	Name          string
	User          string
	Password      string
	EnableSSlMode bool
}

type Config struct {
	Version      string
	ServiceName  string
	HttpPort     int
	JwtSecretKey string
	DB           *DBConfig
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load the env variables:", err)
		os.Exit(1)
	}

	VERSION := os.Getenv("VERSION")
	if VERSION == "" {
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

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey == "" {
		fmt.Println("Jwt secret key is required")
		os.Exit(1)
	}

	dbHost := os.Getenv("HOST")
	if dbHost == "" {
		fmt.Println("DB Host is required")
		os.Exit(1)
	}

	dbPort := os.Getenv("PORT")
	if dbPort == "" {
		fmt.Println("DB port is required")
		os.Exit(1)
	}
	DbPort, err := strconv.ParseInt(dbPort, 10, 64)
	if err != nil {
		fmt.Println("Port must be number")
		os.Exit(1)
	}

	dbName := os.Getenv("NAME")
	if dbName == "" {
		fmt.Println("DB Name is required")
		os.Exit(1)
	}

	dbUser := os.Getenv("USER")
	if dbUser == "" {
		fmt.Println("DB User is required")
		os.Exit(1)
	}

	dbPassword := os.Getenv("PASSWORD")
	if dbPassword == "" {
		fmt.Println("DB Password is required")
		os.Exit(1)
	}

	enableSSLMode := os.Getenv("ENABLE_SSL_MODE")

	parsedEnableSSLMode, err := strconv.ParseBool(enableSSLMode)
	if err != nil {
		fmt.Println("Invalid Enable SSL MODE value", err)
		os.Exit(1)
	}

	dbConfig := &DBConfig{
		Host: dbHost,
		Port: int(DbPort),
		Name: dbName,
		User: dbUser,
		Password: dbPassword,
		EnableSSlMode: parsedEnableSSLMode,
	}

	configurations = &Config{
		Version:      VERSION,
		ServiceName:  SERVICE_NAME,
		HttpPort:     int(port), // typecasting
		JwtSecretKey: jwtSecretKey,
		DB: dbConfig,
	}
}

func GetConfig() *Config {
	if configurations == nil {
		loadConfig()
	}
	return configurations
}
