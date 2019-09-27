package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		if i > 5 {
			break //loop is terminated if i > 5
		}
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\nline after for loop\n")
	n := 5
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	fmt.Println("-----------------------------------")

	// We try to break when i and j are the same value
	for i := 0; i < 3; i++ {
		for j := 1; j < 4; j++ {
			fmt.Printf("i = %d , j = %d\n", i, j)
			if i == j {
				break
			}
		}

	}

	fmt.Println("-----------------------------------")

	// Since the above only breaks the inner loop, we use a label to break free ;)
outer:
	for i := 0; i < 3; i++ {
		for j := 1; j < 4; j++ {
			fmt.Printf("i = %d , j = %d\n", i, j)
			if i == j {
				break outer
			}
		}

	}

	fmt.Println("-----------------------------------")
	fmt.Println("This here:")
	for i := 0; i <= 10; i += 2{ // initialisation and post are omitted
		fmt.Printf("%d ", i)
	}
	fmt.Println("\nis the same as: ")
	index := 0
	for ; index <= 10; { // initialisation and post are omitted
		fmt.Printf("%d ", index)
		index += 2
	}
	fmt.Println("\nis the same as: ")
	index = 0
	for index <= 10 { // semicolons are omitted and only condition is present
		fmt.Printf("%d ", index)
		index += 2
	}

	fmt.Println("\n-----------------------------------")
	for no, i := 10, 1; i <= 10 && no <= 19; i, no = i+1, no+1 { //multiple initialisation and increment
		fmt.Printf("%d * %d = %d\n", no, i, no*i)
	}

}
