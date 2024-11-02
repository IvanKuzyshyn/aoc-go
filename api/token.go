package api

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const (
	token = "SESSION_TOKEN"
)

func getSessionToken() (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", err
	}
	token := os.Getenv(token)
	if len(token) == 0 {
		return "", fmt.Errorf("value of %s is not set or is empty", token)
	}
	return token, err
}
