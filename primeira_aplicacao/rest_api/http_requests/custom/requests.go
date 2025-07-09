package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	req, err := http.NewRequest(http.MethodGet, "https://www.google.com", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("ACCEPT", "application/json")
	req.Header.Set("authorization", "123")
	req.Header.Add("", "")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	createContext()

	fmt.Println(string(data))
}

func createContext() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5 * time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		"https://www.google.com",
		nil,
	)
	if err != nil {
		panic(err)
	}

	req.Header.Set("ACCEPT", "application/json")
	req.Header.Set("authorization", "123")
	req.Header.Add("", "")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
