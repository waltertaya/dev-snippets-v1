package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func main() {
	apiURLs := []string{
		"https://dummyjson.com/users",
		"https://dummyjson.com/products",
		"https://dummyjson.com/carts",
		"https://dummyjson.com/todos",
		"https://dummyjson.com/posts",
	}

	results := make(chan string)
	var wg sync.WaitGroup

	for _, url := range apiURLs {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			result := fetchData(url)
			results <- result
		}(url)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println(result)
	}
}

func fetchData(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		return "Failed to fetch data from " + url + ": " + err.Error()
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "Failed to read response from " + url + ": " + err.Error()
	}

	return fmt.Sprintf("Data fetched successfully from %s\nResponse: %s", url, string(body[:100])) // Only print first 100 chars
}
