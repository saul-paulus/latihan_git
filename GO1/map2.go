// to take input from the user
package main

import "fmt"

type Person struct {
	Name  string
	Likes []string
}

// main function
func main() {
	var people []*Person

	likes := make(map[string][]*Person)
	for _, p := range people {
		for _, l := range p.Likes {
			likes[l] = append(likes[l], p)
		}
	}

	for _, p := range likes["cheese"] {
		fmt.Println(p.Name, "likes cheese.")
	}

	fmt.Println(len(likes["bacon"]), "people like bacon.")
}
