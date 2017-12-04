package main

import "fmt"

func duck(k ...int) int {
	var greatest int
	for _, i := range k {
		if i > greatest {
			greatest = i
		}
	}
	return greatest
}

func main() {
	a := []int{3, 996, 9, 8, 7, 56, 41, 23, 98, 54, 74, 23, 65, 74, 15, 26, 35, 96, 81}
	g := duck(a...)
	fmt.Println(g)
}
