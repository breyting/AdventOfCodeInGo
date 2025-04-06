package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	req, err := http.NewRequest("GET", "https://adventofcode.com/2015/day/2/input")
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-type", "text/plain")

	if res.StatusCode != 200 {
		fmt.Println("Not the good status code :", res.StatusCode)
	}

}
