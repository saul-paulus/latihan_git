package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("==================FACTORIAL==================")
	// Println function is used to
	// display output in the next line
	fmt.Println("Masukkan Nilai Faktorial: ")

	// var then variable nilai then variable type
	var nilai uint64

	// Taking input from user
	fmt.Scanln(&(nilai))

	start := time.Now()
	for i := uint64(1); i <= nilai; i++ {
		fmt.Printf("Factorial for %d is : %d \n", i, Factorial(uint64(i)))
	}
	end := time.Now()
	fmt.Printf("Calculation finished in %s \n", end.Sub(start)) //Calculation finished in 2.0002ms
}

func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}
