// to take input from the user
package main

import "fmt"

type Kendaraan struct {
	jenis string
	warna string
}

// main function
func main() {
	var m = 1000
	var x [1000]Kendaraan

	for j := 0; j < m; j++ {
		var kendaraan Kendaraan
		kendaraan.jenis = fmt.Sprint("jenis", j)
		kendaraan.warna = fmt.Sprint("warna", j)
		x[j] = kendaraan
	}

	for k := 0; k < m; k++ {
		fmt.Printf("Jenis : %s, Warna : %s\n", x[k].jenis, x[k].warna)
	}

}
