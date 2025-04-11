package main

import "fmt"

func isPrime (number int) bool { 
    if (number < 2){
		return false;
	}
    if (number == 2){ return true;}
    if (number % 2 == 0) {return false;}
    for i:=3; (i*i) <= number; i+=2 {
        if (number % i == 0 ) {return false;}
    }
    return true;
}

func findPrimeNumbers(n int) {
	// constraint n > 2
	arePrime := []int{} // slice to store prime

	i := 2

	for i <= n {
		result := isPrime(i)
		if result {
			arePrime = append(arePrime, i)
		}
		i++
	}
	fmt.Println(arePrime)
}

func main() {
	var n int
	fmt.Print("Enter the end number to find prime numbers: ")
	fmt.Scan(&n)

	findPrimeNumbers(n)
}
