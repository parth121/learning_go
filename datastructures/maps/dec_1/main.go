package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	m["c"] = 3
	delete(m, "b")
	a, b := m["c"]
	fmt.Println(len(m))
	fmt.Println(m, a, b)
}
