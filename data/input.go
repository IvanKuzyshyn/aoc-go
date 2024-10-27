package data

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	originUrl = "https://adventofcode.com"
)

type Input struct {
	Day  int
	Year int
}

func (i *Input) Load() ([]byte, error) {
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
