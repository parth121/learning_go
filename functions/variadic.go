package main

import "fmt"

func main() {
	data := []float64{2, 4}
	fmt.Println(average(data...))
}

func average(nm ...float64) float64 {
	var total float64
	for _, i := range nm {
		total += i
	}
	return total / float64(len(nm))
}
