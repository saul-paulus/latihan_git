// to take input from the user
package main

import "fmt"

type Vertex struct {
	Lat,
	Long float64
}

// main function
func main() {
	m := make(map[string]Vertex)
	m["Medan"] = Vertex{98.67123, 3.66681}
	m["Banda Aceh"] = Vertex{95.34312, 5.54491}
	m["Denpasar"] = Vertex{115.224609, -8.643480}
	m["Bengkulu"] = Vertex{102.304688, -3.794373}
	m["Yogyakarta"] = Vertex{110.37625, -7.80279}
	m["Jakarta Pusat"] = Vertex{106.826591, -6.176396}
	fmt.Println(m["Banda Aceh"])
}
