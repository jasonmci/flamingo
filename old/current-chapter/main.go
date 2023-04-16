package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	} 
	return n * factorial(n - 1)
}

func main() {
	n := 8
	fmt.Printf("Factorial of %d is %d", n, factorial(n))
}