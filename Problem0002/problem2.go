package main

import "fmt"

func main() {
	n2 := 1
	n1 := 2
	n := 0
	result := 2 // Result starts from 2 because 2nd Fibonacci number is even.

	for n < 4000000 {
		n = n1 + n2
		n2 = n1
		n1 = n

		if n%2 == 0 {
			result += n
		}
	}

	fmt.Printf("Result is %d\n", result)
}
