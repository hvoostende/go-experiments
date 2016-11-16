package test

import "fmt"

//Pr prints without fmt
func Pr(str interface{}) {
	fmt.Println(str)
}

//Inc Gives a new number
func Inc(i int) func() int {
	x := i
	return func() int {
		x++
		return x

	}
}
