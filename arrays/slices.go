package main

import (
	"fmt"
)

func main() {

	{
		a := [...]int{76, 77, 78, 79, 80}
		var b []int = a[1:4] //creates a slice from a[1] to a[3]
		fmt.Println(b)
	}
	fmt.Println("-----------------------------------")
	{
		a := [...]int{76, 77, 78, 79, 80}
		var b = a[1:4]
		fmt.Println(b)
	}
	fmt.Println("-----------------------------------")
	{
		a := []int{76, 77, 78, 79, 80}
		fmt.Println(a)
	}
	fmt.Println("-----------------------------------")
	{
		arr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
		slice := arr[2:5]
		fmt.Println("array before change", arr)
		for i := range slice {
			slice[i]++
		}
		fmt.Println("array after change", arr)
		// Note how the indexes 2, 3, 4 changed
	}
	fmt.Println("-----------------------------------")
	{
		numa := [3]int{78, 79 ,80}
		nums1 := numa[:] //creates a slice which contains all elements of the array
		nums2 := numa[:]
		fmt.Println("array before change 1",numa)
		nums1[0] = 100
		fmt.Println("array after modification to slice nums1", numa)
		nums2[1] = 101
		fmt.Println("array after modification to slice nums2", numa)
	}
}
