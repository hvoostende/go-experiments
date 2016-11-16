package main

import (
	"fmt"
	"math"
)

func main() {
	var inputNumber float64
	var inputApproximate float64

	fmt.Print("Calculate square root of number: ")
	fmt.Scan(&inputNumber)
	fmt.Print("Approximate is reached with (for example 0.001): ")
	fmt.Scan(&inputApproximate)

	output, total := newton(inputNumber, inputApproximate)

	fmt.Printf("The square root using Math.Sqrt: %v\n", math.Sqrt(inputNumber))
	fmt.Printf("The square root using Newton's Method: %v\n", output)
	fmt.Printf("It took %v run's to get to an approximate of %v\n", total, inputApproximate)

}

//Newton calculates the squareroot of a number using Newon's Method
func newton(x float64, a float64) (float64, int) {

	if a == 0 {
		return 0, 0

	}
	if x == 0 {
		return 0, 0
	}

	if x == 1 {
		return 1, 0
	}

	z := 1.0
	var i int

	for {
		i++
		z = z - ((math.Pow(z, 2) - x) / (2 * x))
		if math.Sqrt(x)-z < a {
			break
		}
	}
	return z, i

}
