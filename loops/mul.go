package main

import "fmt"

func main() {
	var sum = 0
	for i := 0; i < 10; i++ {
		if i%3 == 0 {
			sum = sum + i
		} else if i%5 == 0 {
			sum = sum + i
		}
	}
	fmt.Println(sum)
}
