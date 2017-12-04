package main

import "fmt"

func main() {
	var (
		a = 10
		b = 2.3
		c = "hello pussy"
		d = true
	)

	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
}
