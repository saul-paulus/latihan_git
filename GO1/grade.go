// to take input from the user 
package main 

import "fmt"

// main function 
func main() { 

	// Println function is used to 
	// display output in the next line 
	fmt.Println("Masukkan Nilai N: ") 

	// var then variable nilai then variable type 
	var nilai uint64

	// Taking input from user 
	fmt.Scanln(&(nilai)) 	 

	// Print function is used to 
	// display output in the same line 
	fmt.Print("Grade Nilai Anda adalah : ") 
    //A (85 – 100), B (75 – 84), C (60 – 74), D(0 – 59), sisanya salah format
    switch {
    case (nilai) >= 85 && (nilai) <= 100:
        fmt.Println("A")
    case (nilai) >= 75 && (nilai) <= 84:
        fmt.Println("B")
    case (nilai) >= 60 && (nilai) <= 74:
        fmt.Println("C")
    case (nilai) >= 0 && (nilai) <= 59:
        fmt.Println("D")
    default:
        fmt.Println("Salah Format")
    }	
} 
