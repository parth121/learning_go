package main

import "fmt"

func fun(x *int) {
	*x = 0
}

var x int = 47

func main() {
	fun(&x)
	fmt.Println(x)
}
