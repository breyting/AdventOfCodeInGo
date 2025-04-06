package config

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const baseURL = "https://adventofcode.com/%s/day/%s/input"

func ReadInputAOC(year, day string) (string, error) {
	url := fmt.Sprintf(baseURL, year, day)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error creating HTTP request: %v\n", err)
		return "", err
	}

	// Use session token from environment variable
	req.Header.Add("Cookie", "session=ru=53616c7465645f5f459bfaca53ef1449f55a34ccc459db99fd680e614846e262b230b453dbfabd0b18caf8096ade16cd; session=53616c7465645f5f88bf662419ef7a32b9bedaa5f086018dbae8279421398dcbfded47206a81fc3255c75d08ea4ae0da023e042e65951321bf59b426316da8b3"+os.Getenv("AOC_SESSION"))

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
