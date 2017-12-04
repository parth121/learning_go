package main

import "fmt"

func main() {
	for i := 200; i < 500; i++ {
		fmt.Printf("%d \t - %q\n", i, i)
	}
}
