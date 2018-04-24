## iKariera Tasks

This project is about a Command Line Tool to solve mathematical problems using Golang. There are two following problems.
  * Calculate Pi for given n number of terms
  * Frequency of digits in all Prime Numbers

```
git clone https://github.com/sbalaji05/iKarieraTasks.git
```

### Install Go
Install Golang for Ubuntu using the following command
```
sudo apt-get install golang-go
```

### Usage

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

### CalculatePi
This problem is to calculate pi value based on the number of terms n given. The default value of n is taken as 99999 if not provided any input. The higher the n value, more accurate the pi value. </br>
Eg: a(10) = 3.0418396 </br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;a(99999) = 3.1416059

### Frequency of digits in all Prime Numbers
For a given integer limit, this problem finds all the prime numbers and finds the Frequency of digits in all prime numbers. The default value is 1000 if the input is not provided. 

### Alternate way
  If Golang is not supported in your system or facing some issues in running the project, the same mathematical tasks are implemented in simple C++ programs. The c++ programs are available under C++ folder.
