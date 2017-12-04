package main

import "fmt"

func tada() func() string {
	return func() string {
		return fmt.Sprint("Go is Great")
	}
}

func main() {
	yo := tada()
	fmt.Printf("%T\n", yo)
	fmt.Println(yo())
}
