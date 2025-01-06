package config

import (
	"log"
	"os"
	"strconv"
)

func GetEnv() string {
	return getEnvironmentValue("ENV")
}

func getEnvironmentValue(key string) string {
	if os.Getenv(key) == "" {
		log.Fatalf("%s environment variable is missing.", key)
	}
	return os.Getenv(key)
}

func GetApplicationHttpPort() int {
	portStr := getEnvironmentValue("APPLICATION_HTTP_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("port: %s is invalid", portStr)
	}
	return port
}
