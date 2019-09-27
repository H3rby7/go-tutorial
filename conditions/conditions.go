package main

import (
	"fmt"
)

func main() {
	num := 10
	if num%2 == 0 {
		//checks if number is even
		fmt.Println("the number is even")
	} else {
		fmt.Println("the number is odd")
	}
	if num2 := 21; num2%2 == 0 {
		fmt.Println("the second number is even")
	} else if num2%2 == 1 {
		fmt.Println("the second number is uneven")
	} else {
		fmt.Println("You broke Math")
	}
	// Not the var "num2" is only available to the conditional statement and the if/else blocks
	fmt.Println("Number one is", num, "Number two is 21")

	finger := 4
	switch finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")
	default:
		fmt.Println("not a finger for most humans")
	}

	switch finger := 2; finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")
	default:
		fmt.Println("not a finger for most humans")
	}

	switch letter := "i"; letter {
	case "a", "e", "i", "o", "u": //multiple expressions in case
		fmt.Println("vowel")
	default:
		fmt.Println("not a vowel")
	}

	number := 75
	// expression is omitted, acts like multiple if else's
	switch {
	case number >= 0 && number <= 50:
		fmt.Println("number is greater than 0 and less than 50")
	case number >= 51:
		// Wil jump into this line of code and never reach the last statement
		fmt.Println("number is greater than 51")
	case number >= 71:
		fmt.Println("number is greater than 71")
	default:
		fmt.Println("I am a default statement")
	}


	fmt.Println("\n-----------------------------------")
	// fallthrough is like omitting java's BREAK
	switch number := 75; {
	case number >= 0 && number <= 50:
		fmt.Println("number is greater than 0 and less than 50")
		fallthrough
	case number >= 51:
		// Wil jump into this line of code and never reach the last statement
		fmt.Println("number is greater than 51")
		fallthrough
	case number >= 71:
		fmt.Println("number is greater than 71")
		fallthrough
	default:
		// Note that we even get to this statement using fallthrough!
		fmt.Println("I am a default statement")
	}
}
