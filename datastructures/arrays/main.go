package main

import "fmt"

func main() {
	var x [10]int
	fmt.Println(x)
	for i := 0; i < 10; i++ {
		x[i] = i
	}
	fmt.Println(cap(x))
	fmt.Println(x)
}
