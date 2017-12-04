package main

import "fmt"

func main() {
	var a int
	fmt.Print("Enter an integer: ")
	fmt.Scan(&a)
	fmt.Println(half(a))
}

func half(m int) (int, bool) {
	x := m / 2
	var typ bool
	if m%2 == 0 {
		typ = true
	} else {
		typ = false
	}
	return x, typ
}
