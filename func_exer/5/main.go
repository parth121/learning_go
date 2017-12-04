package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	switch a {
	case 1:
		foo(1, 2)
	case 2:
		foo(1, 2, 3)
	case 3:
		aSlice := []int{1, 2, 3, 4}
		foo(aSlice...)
	default:
		foo()
	}
}

func foo(xxx ...int) {
	for _, i := range xxx {
		fmt.Println(i)
	}
}
