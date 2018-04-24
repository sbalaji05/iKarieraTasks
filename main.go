package main

import (
	"flag"
	"fmt"
)

// Main function
func main() {
	var option, n, primelimit int
	var ErrorMsg string
	// By defualt the program takes input as 1 which is CalculatePi problem
	flag.IntVar(&option, "option", 1, "Provide 1 for CalculatePi or \n\tProvide 2 for Frequency of digits in Prime Numbers.")
	// By default the n of terms is taken as 99999 if not provided manually
	flag.IntVar(&n, "n", 99999, "Enter number of terms (n) to calculate pi.")
	// By default prime numbers are found from 1 to 1000 if not provided
	flag.IntVar(&primelimit, "primelimit", 1000, "Enter the limit to find Prime Numbers.")
	flag.Parse()

	if option < 0 {
		fmt.Println("\nOption should not be a negative number. \nPlease provide either 1 or 2!\n")
		flag.Usage()
	} else if option == 1 { // CalculatePi problem
		ErrorMsg = Validate(n)
		if ErrorMsg != "" {
			fmt.Println("\nNumber of terms n", ErrorMsg)
			flag.Usage()
		} else {
			CalculatePi(n)
		}
	} else if option == 2 { // Frequency of digits in Prime Numbers problem
		ErrorMsg = Validate(primelimit)
		if ErrorMsg != "" {
			fmt.Println("\nPrime Limit", ErrorMsg)
			flag.Usage()
		} else {
			Prime_Numbers(primelimit)
		}
	}
}
