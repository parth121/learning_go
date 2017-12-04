package main

import "fmt"

func main() {
	m := make([]int, 2, 2)
	fmt.Println(cap(m))
	m = append(m, 2, 4, 5)
	fmt.Println(cap(m))
	fmt.Println(len(m))
	fmt.Println(m)
}
