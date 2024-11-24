package data

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const (
	originUrl = "https://adventofcode.com"
	token     = "SESSION_TOKEN"
)

func (i *PuzzleData) GetInput() ([]byte, error) {
	url := fmt.Sprintf("%s/%d/day/%d/input", originUrl, i.Year, i.Day)
	token, err := getSessionToken()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Cookie", fmt.Sprintf("session=%s", token))
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Error closing response body: %v", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed with status code: %d", resp.StatusCode)
	}

	contentType := resp.Header.Get("Content-Type")
	if contentType != "text/plain" {
		return nil, fmt.Errorf("unexpected content type: %s", contentType)
	}

	return io.ReadAll(resp.Body)
}

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
