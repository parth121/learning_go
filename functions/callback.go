package main

import "fmt"

func lele(num []int, test func(int)) {
	for _, i := range num {
		test(i)
	}
}

func main() {
	lele([]int{1, 2, 3, 4, 5}, func(n int) {
		fmt.Println(n)
	})
}
