package main

import "fmt"

func main() {
	a, _ := greet("Rohan", "Raj")
	fmt.Println(a)
}

func greet(fname, lname string) (string, string) {
	return fmt.Sprint(fname, " ", lname), fmt.Sprint("Hello ", "Go")
}
