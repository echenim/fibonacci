package main

import "fmt"

func main() {
	for i := 0; i <= 9; i++ {
		fmt.Printf(" %v ", Fibonacci(i))
	}
}

//Fibonacci function to generate fibonacci sequence
func Fibonacci(num int) int {
	if num <= 1 {
		return num
	}
	return Fibonacci(num-1) + Fibonacci(num-2)
}
