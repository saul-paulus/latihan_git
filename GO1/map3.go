// to take input from the user
package main

import "fmt"

type Vertex struct {
	AreaId,
	Lat,
	Long float64
}

// main function
func main() {
	m := make(map[string]Vertex)
	m["Gorontalo"] = Vertex{501197, 123.08807, 0.53175}
	m["Banda Aceh"] = Vertex{501397, 95.34312, 5.54491}
	m["Denpasar"] = Vertex{501164, 115.224609, -8.643480}
	m["Bengkulu"] = Vertex{501178, 102.304688, -3.794373}
	m["Yogyakarta"] = Vertex{501190, 110.37625, -7.80279}
	m["Jakarta Pusat"] = Vertex{501195, 106.826591, -6.176396}
	fmt.Println(m["Banda Aceh"])
}
