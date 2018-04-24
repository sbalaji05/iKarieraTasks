
#include <stdio.h>
#include <stdlib.h>
#include <iostream>
using namespace std;

int digit_Array[10];

// Function to find frequency of a prime number
int * GetFrequency(int num) {
  int i, j, ctr, r;
  for (i = 0; i < 10; i++) {
    ctr = 0;
    for (j = num; j > 0; j = j / 10) {
      r = j % 10;
      if (r == i) {
        ctr++;
      }
    }
    digit_Array[i] += ctr;
  }
  return digit_Array;
}

int main() {
  int count = 1, * freq_Array;
  cout << "The Prime numbers from 1 to 1000 is " << endl;
  cout << "----------------------------------------" << endl;
  for (int i = 2; i < 1000; i++) // Find prime number till 1000
  {
    bool prime = true;
    for (int j = 2; j < i; j++) {
      if (i % j == 0) {
        prime = false;
        break; // i is NOT a prime number.
      }
    }
    if (prime) {
      cout << i << " ";
      if (count % 10 == 0) // BReak to next line after 10 numbers.
      {
        cout << endl;
      }
      count++;
      freq_Array = GetFrequency(i);
    }
  }
  cout << endl << endl;
  for (int i = 0; i < 10; i++) {
    cout << "The frequency of " << i << " in all prime numbers = " << freq_Array[i] << endl;
  }
  cout << endl;
  return 0;
}
