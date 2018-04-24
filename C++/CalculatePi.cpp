
#include <stdio.h>
#include <stdlib.h>
#include <iostream>
using namespace std;

// Function to validate n
string Validate(int n) {
  string isError;
  if (n < 0) {
    isError = "ERROR! Invalid Input. Please provide positive integer!";
  } else if (n == 0) {
    isError = "ERROR! Zero wouldnt work, Please provide positive integer!";
  }
  return isError;
}

// Function to calculate pi
void CalculatePi(int n) {
  cout.precision(10);
  double sum(0.0), pi;
  //processing
  for (int i = 0; i < n; i++) {
    if (i % 2 == 0)
      sum += 1 / (2.0 * i + 1);
    else
      sum -= 1 / (2.0 * i + 1);
    pi = 4 * sum;
  }
  cout << "a(" << n << ") = " << pi << endl;
  cout << "Approximate value of pi for " << n << " number of terms." << endl;
}

int main() {
  char run;
  do {
    int n;
    string isError;
    //input
    cout << "Calculating pi value based on n terms." << endl;
    cout << "Higher the n, more accurate the pi value Eg:99999" << endl;
    cout << "---------------------------------------------------" << endl << endl;
    cout << "Enter number of terms (n) to calculate pi: ";
    cin >> n;
    isError = Validate(n);
    if (isError == "") {
      CalculatePi(n);
    } else {
      cout << isError << endl;
    }
    cout << "---------------------------------------------------" << endl << endl;
    cout << "Run program again for different n? Press 'y' or 'Y':";
    cin >> run;
  } while (run == 'y' || run == 'Y');
}
