package main

import (
	"fmt"
)

// Function to calculate pi based on the number of terms n
func CalculatePi(n int) {
	var sum, pi float32
	fmt.Println("\nCalculating pi value based on n terms.")
	fmt.Println("Higher the n, more accurate the pi value Eg:99999")
	fmt.Println("---------------------------------------------------\n")
	//processing
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			sum += 1 / float32((2.0*i + 1))
		} else {
			sum -= 1 / float32((2.0*i + 1))
		}
		pi = 4 * sum
	}
	fmt.Print("a(", n, ") = ", pi)
	fmt.Print("\nApproximate value of pi for ", n, " number of terms.")
	fmt.Println("\n\n---------------------------------------------------")
}
