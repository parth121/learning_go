package main

import "fmt"

func main() {
	for i := 67500; i < 67900; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}
}
