package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a URL as an argument.")
	}

	url := os.Args[1]

	if !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	filename := "result.json"
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Write response to file
	n, err := io.Copy(f, res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Wrote %d bytes to %s\n", n, filename)
	fmt.Printf("HTTP Status: %s\n", res.Status)
}
