package unmarshalling

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func Unmarshal() {
	// Sample JSON
	jsonString := `{"name":"Walter Onyango","age":21,"email":"walter@example.com"}`

	var user User

	// Unmarshal the JSON into the struct
	err := json.Unmarshal([]byte(jsonString), &user)
	if err != nil {
		fmt.Println("Error unmarshalling:", err)
		return
	}

	// Access the struct fields
	fmt.Println("Name:", user.Name)
	fmt.Println("Age:", user.Age)
	fmt.Println("Email:", user.Email)
}
