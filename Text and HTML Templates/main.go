package main

import (
	"html/template"
	"net/http"
)

type User struct {
	Name string
	Age  int
}

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template.html"))

	user := User{Name: "Walter", Age: 21}
	tmpl.Execute(w, user)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
