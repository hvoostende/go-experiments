package main

import "fmt"

func main() {
	number := 26
	c := incrementor(number)
	cfactorial := factorial(c)
	fmt.Println("number:", number)
	for n := range cfactorial {
		fmt.Println("1:", n)
	}
	fmt.Println("2:", f)
}

func incrementor(n int) chan int {
	out := make(chan int)
	go func() {
		for i := n; i > 0; i-- {
			out <- i
		}
		close(out)
	}()
	return out
}

func factorial(c chan int) chan int {
	out := make(chan int)
	go func() {
		product := 1
		for n := range c {
			product *= n
		}
		out <- product
		close(out)
	}()
	return out
}

func factorial2(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}
