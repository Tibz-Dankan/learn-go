// Question 1: Calculate Factorial
// Write a Go program that calculates the factorial
// of a given non-negative integer. The program should
// prompt the user to enter an integer and then calculate
// its factorial. Ensure that the input is validated
// to be a non-negative integer.

// Solution

package main

import (
	"fmt"
)

// func main() {
	func facto() {
	var num int
	factorial := 1

	fmt.Println("Enter an integer number:")
	_, err := fmt.Scanln(&num)

	if err != nil || num <= 0 {
		fmt.Println("Invalid input. Number must be an integer greater than 0.")
		return
	}

	for i := num; i > 0; i-- {
		factorial *= i
	}

	fmt.Printf("Factorial of %d is: %d\n", num, factorial)
}

