package main

import (
	"fmt"
)

func main() {
	var a [3]int //int array with length 3
	fmt.Println(a)
	fmt.Println("-----------------------------------")
	a[0] = 12 // array index starts at 0
	a[1] = 78
	a[2] = 50
	fmt.Println(a)
	fmt.Println("-----------------------------------")
	a2 := [3]int{12, 78, 50} // short hand declaration to create array
	fmt.Println(a2)
	fmt.Println("-----------------------------------")
	fmt.Println([3]int{12, 78, 50})
	fmt.Println("-----------------------------------")
	// Only the first element is 12, the rest is still 0s
	fmt.Println([3]int{12})
	fmt.Println("-----------------------------------")
	// Compiler determines the length of the array
	fmt.Println([...]int{12, 78, 50})
	fmt.Println("-----------------------------------")

	// The size of an array is part of the type, so the below stuff does not work.
	{
		// a := [3]int{5, 78, 8}
		// var b [5]int
		// b = a
	}

	// Arrays are value types. So they do not get passed by reference.
	{
		a := [...]string{"USA", "China", "India", "Germany", "France"}
		b := a // a copy of a is assigned to b
		b[0] = "Singapore"
		fmt.Println("a is ", a)
		fmt.Println("b is ", b)
	}

	fmt.Println("-----------------------------------")

	{
		num := [...]int{5, 6, 7, 8, 8}
		fmt.Println("before passing to function ", num)
		changeLocal(num) //num is passed by value
		fmt.Println("after passing to function ", num)
		// array in this scope is the same before and after the function call
		// By the way its length is
		fmt.Println("-----------------------------------")
		fmt.Println("Length is", len(num))

		fmt.Println("-----------------------------------")
		for i := 0; i < len(num); i++ {
			fmt.Printf("%d th element of a is %.2f\n", i, float64(num[i]))
		}

		fmt.Println("-----------------------------------")
		sum := 0
		for i, v := range num {//range returns both the index and value
			fmt.Printf("%.2d the element of a is %.3d\n", i, v)
			sum += v
		}
		fmt.Println("\nsum of all elements of a",sum)
	}
}


func changeLocal(num [5]int) {
	num[0] = 55
	fmt.Println("inside function ", num)

}