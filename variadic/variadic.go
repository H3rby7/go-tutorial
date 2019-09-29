package main

import "fmt"

func main() {
	{
		fmt.Println("Sum of 4,89,1 is", addUp(4, 89, 1))
		mySlice := []int{69, 42, 1337}
		// Note that we can pass the slice to a variadic function by appending '...'
		// When we do this, our slice will be used as argument to the function. So no new slice is created.
		// Any modification we make to the slice inside the function will affect the original slice.
		fmt.Println(mySlice, "added up is:", addUp(0, mySlice...))
	}
	fmt.Println("\n-----------------------------------")
	{
		texts := []string{"Go", "is", "King"}
		result := sliceFun(texts...)
		fmt.Println(result)
	}
}

func addUp(a int, b ... int) int {
	fmt.Printf("type of b is %T\n", b)
	sum := a
	for _, v := range b {
		sum += v
	}
	return sum
}

func sliceFun(input ... string) []string {
	input[0] = "Joka"
	return input
}
