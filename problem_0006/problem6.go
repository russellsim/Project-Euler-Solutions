package main

import "fmt"

/*
Notice that:
(a + b)^2 = a^2 + b^2 + 2(ab)
(a + b + c)^2 = a^2 + b^2 + c^2 + 2(ab + ac + bc)
(a + b + c + d)^2 = a^2 + b^2 + c^2 + d^2 + 2(ab + ac + ad + ... (Rest of 2 letter combinations of abcd))

Therefore the difference is just 2(ab + ac + ad + ... (Rest of 2 letter combinations of abcd)
*/

func main() {
	limit := 100
	result := 0
	iterations := 0

	for i := 1; i <= limit; i++ {
		for j := 1; j <= limit; j++ {
			if i < j {
				result += 2 * i * j
				iterations++
			}
		}
	}

	fmt.Println("Result: ", result)
	fmt.Println("Iterations: ", iterations)
}
