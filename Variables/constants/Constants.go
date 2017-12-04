package main

import "fmt"

const p = "This is some demo text"
const q = 47

func main() {
	const (
		A = iota
		B = iota
		C = iota
	)
	fmt.Println(p)
	fmt.Println(q)
	fmt.Println(A, B, C)
}
