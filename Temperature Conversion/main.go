package main

import (
    "fmt"
    "temperature-conversion/temp"
)

func main() {
    results := temp.CToF(92)
    fmt.Println(results)

    results = temp.Fahrenheit(123)
    fmt.Println(results)
}
