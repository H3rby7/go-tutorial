package main

import (
	"fmt"
	"reflect"
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
		numa := [3]int{78, 79, 80}
		nums1 := numa[:] //creates a slice which contains all elements of the array
		nums2 := numa[:]
		fmt.Println("array before change 1", numa)
		nums1[0] = 100
		fmt.Println("array after modification to slice nums1", numa)
		nums2[1] = 101
		fmt.Println("array after modification to slice nums2", numa)
		changeE2(nums1)
		fmt.Println("array after modification via function", numa)
	}
	fmt.Println("-----------------------------------")
	{
		fruitarray := []string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
		fruitslice := fruitarray[1:3]
		fmt.Printf("fruitarray is of type: %T and fruitslice is of type %T\n", fruitarray, fruitslice)
		//length of fruitslice is 2 and capacity is 6 (capacity is 6, because the original array is 7 pieces long and we started at index 1)
		fmt.Printf("length of slice %d capacity %d\n", len(fruitslice), cap(fruitslice))
		//re-slicing fruitslice till its capacity
		fruitslice = fruitslice[:cap(fruitslice)]
		fmt.Println("After re-slicing length is", len(fruitslice), "and capacity is", cap(fruitslice))
	}
	fmt.Println("-----------------------------------")
	{
		// Creates an array and a slice reference to it
		i := make([]int, 5, 5)
		i[2] = 3
		fmt.Println(i)
	}
	fmt.Println("-----------------------------------")
	{
		cars := []string{"Ferrari", "Honda", "Ford"}
		//capacity of cars is 3
		fmt.Println("cars is a ", reflect.TypeOf(cars), ":", cars, "has old length", len(cars), "and capacity", cap(cars))
		cars = append(cars, "Toyota")
		//capacity of cars is doubled to 6, because we append an element that does not fit the capacity
		fmt.Println("cars is a ", reflect.TypeOf(cars), ":", cars, "has new length", len(cars), "and capacity", cap(cars))
	}
	fmt.Println("-----------------------------------")
	{
		//zero value of a slice is nil
		var names []string
		if names == nil {
			fmt.Println("slice is nil going to append")
			names = append(names, "John", "Sebastian", "Vinay")
			fmt.Println("names contents:", names)
		}
	}
	fmt.Println("-----------------------------------")
	{
		veggies := []string{"potatoes", "tomatoes", "brinjal"}
		fruits := []string{"oranges", "apples"}
		// append one slice to another using '...'
		food := append(veggies, fruits...)
		fmt.Println("food:", food)
	}
}

func changeE2(slice []int) {
	slice[2] = 34
}
