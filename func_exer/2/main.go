package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	m := func(a int) (int, bool) {
		return a / 2, a%2 == 0
	}

	fmt.Println(m(a))
}
