package main

import "fmt"

func main() {
	myprint := func(data ...int) {
		for i, j := range data {
			fmt.Println(i, "-", j)
		}
		fmt.Print("\n")
	}

	list := []int{1, 2, 3}
	myprint(list...)
}
