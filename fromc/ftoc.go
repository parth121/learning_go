package main

import "fmt"

const (
	LOWER = 0.0
	UPPER = 300.0
	STEP  = 20.0
)

func main() {
	conv()
	//fmt.Println(conv())
}

func conv() { //[]float64 {
	//op := make([]float64, 100, 100)
	for fahr := LOWER; fahr <= UPPER; fahr = fahr + STEP {
		celcius := (5.0 / 9.0) * (fahr - 32)
		fmt.Printf("%v\t%v\n", fahr, celcius)
		//op = append(op, celcius)
	}
	//return op
}
