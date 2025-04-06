package config

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const baseURL = "https://adventofcode.com/2024/day/%s/input"

func ReadInput(day string) (string, error) {
	url := fmt.Sprintf(baseURL, day)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error creating HTTP request: %v\n", err)
		return "", err
	}

	// Use session token from environment variable
	req.Header.Add("Cookie", "session="+os.Getenv("AOC_SESSION"))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making HTTP request: %v\n", err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Unexpected HTTP status: %d\n", resp.StatusCode)
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return "", err
	}

	return string(body), nil
}
