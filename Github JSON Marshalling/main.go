package main

import (
	"fmt"
	"github/github"
)

func main() {

	result, err := github.SearchIssues([]string{"python", "json", "encoding", "go"})
	if err != nil {
		panic(err)
	}

	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}
