package main

import (
	"fmt"
)

func main() {
	num := 10
	if num % 2 == 0 {
		//checks if number is even
		fmt.Println("the number is even")
	}  else {
		fmt.Println("the number is odd")
	}
	if num2 := 21; num2 % 2 == 0 {
		fmt.Println("the second number is even")
	} else if num2 % 2 == 1{
		fmt.Println("the second number is uneven")
	} else {
		fmt.Println("You broke Math")
	}
	// Not the var "num2" is only available to the conditional statement and the if/else blocks
	fmt.Println("Number one is", num, "Number two is 21")
}
