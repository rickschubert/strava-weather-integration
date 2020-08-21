package utils

import (
	"os"
	"strconv"
)

func GetPortToUse() int {
	defaultPort := 8080
	port := os.Getenv("PORT")
	if port == "" {
		return defaultPort
	}
	portAsInt, err := strconv.Atoi(port)
	if err != nil {
		return defaultPort
	}
	return portAsInt
}
