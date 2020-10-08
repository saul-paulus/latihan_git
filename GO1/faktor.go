// to take input from the user
package main

import "fmt"

// main function
func main() {

	// Println function is used to
	// display output in the next line
	fmt.Println("Masukkan Nilai Faktor: ")

	// var then variable nilai then variable type
	var nilai uint64

	// Taking input from user
	fmt.Scanln(&(nilai))

	// Print function is used to
	// display output in the same line
	for i := uint64(1); i <= nilai; i++ {
		//fmt.Printf("Factorial for %d is : %d \n", i, Factorial(uint64(i)))
		if nilai%i == 0 {
			fmt.Println("Nilai : ", i)
		}
	}
}
