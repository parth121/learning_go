package main

import "fmt"

func main() {
	var test bool
	test = (true && false) || (false && true) || !(false && false)
	fmt.Println(test)
}
