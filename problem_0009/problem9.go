package main

import "fmt"

// Brute force because limit is only 1000
func main() {
	limit := 1000

	for a := 1; a < limit; a++ {
		for b := a + 1; b < limit; b++ {
			for c := b + 1; c < limit; c++ {
				if a*a+b*b == c*c && a+b+c == limit {
					fmt.Println("A B C is: ", a, b, c)
					fmt.Println("Result is: ", a*b*c)
					return
				}
			}
		}
	}
}
