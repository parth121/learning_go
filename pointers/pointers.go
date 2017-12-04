package main

import "fmt"

func main() {
	a := 47
	var b *int = &a
	fmt.Println(&a)
	fmt.Println(*b)
	*b = 62
	fmt.Println(a)
}
