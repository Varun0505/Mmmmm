package main

import "fmt"
import "math"

 	//to get the count of digits
	func DigitsCount(number int) int {
		count := 0
		for number != 0 {
		number /= 10
		count += 1
		}
		return count
	}

	func main() {
		var lastDigit, number int
		var sum float64 = 0
		var count int = 0
		var tempNumber int = 0
	

	fmt.Print("Enter a n digit number : ")
 	fmt.Scanf("%d", &number)//input a number

	tempNumber = number//assigning the number to a temporary variable
	 
	count = DigitsCount(number)//to get the number of digits in the number

 	for {
		// get the right most digit
 		lastDigit = tempNumber % 10
 		sum += math.Pow(float64(lastDigit),float64(count))//add last digit raised to the power of n to sum

 		// update the the temporary number by removing the last digit
 		tempNumber /= 10

 		if tempNumber == 0 {
 			// break statement for stopping the loop
 			break
 		}
 	}

	if number == int(sum) {
 		fmt.Println(number, "is an Armstrong number!")
 	} else {
 		fmt.Println(number, "is NOT an Armstrong number!")
 	}
 }
