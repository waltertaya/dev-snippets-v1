package marshalling

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func Marshal() {
	user := User{
		Name:  "Walter Onyango",
		Age:   21,
		Email: "walter@example.com",
	}

	// Marshal the struct to JSON
	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error marshalling:", err)
		return
	}

	// Print JSON as a string
	fmt.Println(string(jsonData))
}
