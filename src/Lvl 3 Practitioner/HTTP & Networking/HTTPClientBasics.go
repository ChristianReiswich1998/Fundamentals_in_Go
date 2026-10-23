package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Book represents the API response (only fields we need)
type Book struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func GetBookTitle(id int) (string, error) {
	// TODO: Implement this function
	// 1. Build the URL: http://bookstore-api/books/{id}
	// 2. Make a GET request using http.Get()
	// 3. Don't forget to close the response body!
	// 4. Check if status code is http.StatusOK
	// 5. Parse the JSON response into a Book struct
	// 6. Return the title
	URL := fmt.Sprintf("http://bookstore-api/books/%d", id)

	res, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("bad status: %s", res.Status)
	}

	var book Book

	if err := json.Unmarshal(body, &book); err != nil {
		return "", fmt.Errorf("failed to unmarshal json: %w", err)
	}

	return book.Title, nil
}

func main() {
	title, _ := GetBookTitle(1)
	fmt.Println(title)
}
