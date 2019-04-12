package main

import "fmt"

var slice = make([]int, 0, 100)

func isPalindrome(value int) bool {

	slice = slice[:0] // zero the slice

	for value > 0 {
		slice = append(slice, value%10)
		value /= 10
	}

	start := 0
	end := len(slice) - 1

	for start <= end {
		if slice[start] != slice[end] {
			return false
		}
		start++
		end--
	}

	return true
}

func main() {

	result := [3]int{0, 0, 0} // digit1, digit2, result

	for digit1 := 999; digit1 >= 100; digit1-- {
		for digit2 := 999; digit2 >= 100; digit2-- {
			multiply := digit1 * digit2

			if multiply > result[2] {
				if isPalindrome(multiply) {
					result[0] = digit1
					result[1] = digit1
					result[2] = multiply
				}
			}
		}
	}

	fmt.Printf("%d * %d = %d is the largest palindrome from two 3 digit numbers\n", result[0], result[1], result[2])
}
