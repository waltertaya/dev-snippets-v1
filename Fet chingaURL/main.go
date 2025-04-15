package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {

	url := "https://dummyjson.com/products"
	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	b, er := ioutil.ReadAll(res.Body)
	if er != nil {
		log.Fatal(er)
		os.Exit(1)
	}
	fmt.Printf("%T\n", b)

	filename := "data.json"

	e := ioutil.WriteFile(filename, b, 0644)

	if e != nil {
		log.Fatal(e)
		os.Exit(1)
	}
}
