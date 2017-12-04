package main

import "fmt"

var count int64

func main() {
	for fmt.Scan() != EOF {
		count = count + 1
	}
	fmt.Println(count)
}
