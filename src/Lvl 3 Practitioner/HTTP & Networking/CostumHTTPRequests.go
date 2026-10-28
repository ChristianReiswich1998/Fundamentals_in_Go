package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Book2 struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func NewAuthenticatedRequest(method, url, token string) (*http.Request, error) {
	// TODO: Create request with http.NewRequest
	// TODO: Set Authorization header to "Bearer " + token

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	value := fmt.Sprintf("Bearer %s", token)
	req.Header.Set("Authorization", value)

	return req, nil
}

func FetchBookWithAuth(id int, token string) (string, error) {
	// TODO: Build URL for http://bookstore-api/books/{id}
	// TODO: Use NewAuthenticatedRequest to create the request
	// TODO: Execute with client.Do(req)
	// TODO: Check status, parse JSON, return title

	url := fmt.Sprintf("http://bookstore-api/books/%d", id)
	req, err := NewAuthenticatedRequest("GET", url, token)
	if err != nil {
		return "Request failed!", err
	}

	client := &http.Client{Timeout: 10 * time.Second}
	res, err := client.Do(req)
	if err != nil {
		return "Response failed!", err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("bad status: %s", res.Status)
	}

	var book Book2

	if err := json.Unmarshal(body, &book); err != nil {
		return "", fmt.Errorf("failed to unmarshal json: %w", err)
	}

	return book.Title, nil
}

func main() {
	title, err := FetchBookWithAuth(1, "my-token")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(title)
}
