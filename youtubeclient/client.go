package youtubeclient

import (
	"io"
	"log"
	"net/http"
)

func scrap(url string) string {
	// 1. Make the GET request
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	// 2. IMPORTANT: Close the body when the function exits to prevent memory leaks
	defer resp.Body.Close()

	// 3. Read the body content
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// 4. Store in a string variable
	htmlContent := string(bodyBytes)

	// Show a snippet of the result
	return htmlContent
}
