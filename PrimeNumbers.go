package main

import (
	"fmt"
)

var frequency_Array = make([]int, 10)

// Function to find frequency of a prime number
func GetFrequency(num int) []int {
	var i, j, k, count int
	for i = 0; i < 10; i++ {
		count = 0
		for j = num; j > 0; j = j / 10 {
			k = j % 10
			if k == i {
				count++
			}
		}
		frequency_Array[i] += count
	}
	return frequency_Array
}

// Function to find Prime Numbers till given primelimit and to find frequency of digits in all prime numbers
func Prime_Numbers(primelimit int) {
	var count int
	var final_frequency_Array = make([]int, 10)
	count = 1
	fmt.Println("\nThe Prime numbers from 1 to",primelimit,"is ")
	fmt.Println("----------------------------------------\n")

	for i := 2; i < primelimit; i++ { // Find prime numbers till 1000
		isPrime := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				isPrime = false
				break // i is NOT a prime number.
			}
		}
		if isPrime {
			fmt.Print(i, " ")
			if count%10 == 0 { // Break to next line after 10 numbers.
				fmt.Println()
			}
			count++
			final_frequency_Array = GetFrequency(i)
		}
	}
	fmt.Println("\n\n")
	for i := 0; i < 10; i++ {
		fmt.Println("The frequency of ", i, " in all prime numbers = ", final_frequency_Array[i])
	}
}
