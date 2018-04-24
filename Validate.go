package main

// Function to validate input values and return error message if any
func Validate(n int) string {
	var ErrorMsg string
	if n < 0 {
		ErrorMsg = "should not be a negative number. \nPlease provide positve interger!\n"
	} else if n == 0 {
    ErrorMsg = "wouldn't work if its 0. \nPlease provide positive integer!\n"
	} else if n > 2147483647 {
		ErrorMsg = "should not be greater than INT_MAX value. \nPlease provide less than 2147483647.\n"
	}
	return ErrorMsg
}
