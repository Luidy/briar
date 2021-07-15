package config

import (
	"log"
	"os"
	"strconv"
)

type Setting struct {
	ServiceName string
	Host        string
	GRPCPort    string
	HTTPPort    string

	DBName     string
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
}

func NewSetting() Setting {
	return Setting{
		ServiceName: getEnv("service_name", "briar"),
		Host:        getEnv("host", "localhost"),
		GRPCPort:    getEnv("grpc_port", "10620"),
		HTTPPort:    getEnv("http_port", "10621"),

		DBName:     getEnv("db_name", "briardb"),
		DBHost:     getEnv("db_host", "localhost"),
		DBPort:     mustAtoi(getEnv("db_port", "3306")),
		DBUser:     getEnv("db_user", "briaruser"),
		DBPassword: getEnv("db_user", "briarpass"),
	}
}

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}

	return i
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	} else {
		if defaultValue != "" {
			return defaultValue
		}
		log.Fatalf("no env value: %s", key)
	}

	return value
}
