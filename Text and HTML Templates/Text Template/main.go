package main

import (
	"os"
	"text/template"
)

func main() {
	tmpl, _ := template.New("msg").Parse("Hello {{.}}!\n")
	tmpl.Execute(os.Stdout, "Walter")
}
