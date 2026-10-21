package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

// baseURL is the API base URL (can be overridden for testing)
var baseURL = "http://bookstore-api"

// TODO: Implement fetchBook
//
// func fetchBook(ctx context.Context, id int) (string, error)
//
// Steps:
// 1. Build URL: fmt.Sprintf("%s/books/%d", baseURL, id)
// 2. Create request: http.NewRequestWithContext(ctx, "GET", url, nil)
// 3. Execute: http.DefaultClient.Do(req)
// 4. Read body: io.ReadAll(resp.Body)
// 5. Return string(body), nil
//
// Don't forget to import: "context", "fmt", "io", "net/http"
func fetchBook(ctx context.Context, id int) (string, error) {
	var url = fmt.Sprintf("%s/books/%d", baseURL, id)

	request, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return "", err
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", err
	}

	file, _ := io.ReadAll(response.Body)

	return string(file), nil
}

func main() {
}
