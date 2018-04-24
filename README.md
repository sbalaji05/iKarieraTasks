# iKariera Tasks

This project is about a Command Line Tool to solve mathematical problems using Golang. There are two following problems.
  * Calculate Pi for given n number of terms
  * Frequency of digits in all Prime Numbers

```
go get github.concur.com/TravelVendorServices/XSDtoMD
```

## Usage

```
go build
./iKarieraTasks -option=choice of problem -n=Number of terms -primelimit=limit range for prime numbers

Eg: ./iKarieraTasks -option=1 -n=99999
    or
    ./iKarieraTasks -option=2 -primelimit=1000
```
Usage of ./iKarieraTasks:</br>
  &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;-n int </br>
  &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Enter number of terms (n) to calculate pi. (default 99999)</br>
  &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;-option int</br>
  &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Provide 1 for CalculatePi or
	Provide 2 for Frequency of digits in Prime Numbers. (default 1)</br>
  &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;-primelimit int</br>
  &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Enter the limit to find Prime Numbers. (default 1000)
