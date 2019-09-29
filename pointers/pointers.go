package main

import "fmt"

func main() {
	{
		b := 255
		// var 'a' saves the address(*) of an 'int', namely the location(&) of b
		var a *int = &b
		fmt.Printf("Type of a is %T\n", a)
		fmt.Println("address of b is", a)
		// initialized as 'nil'
		var c *int
		fmt.Println("address of c is", c)
	}
	fmt.Println("\n-----------------------------------")
	{
		// Creates a pointer to a new INT
		size := new(int)
		fmt.Printf("Size value is %d, type is %T, address is %v\n", *size, size, size)
		// Access the value of 'size' (term: "dereference") and set it to 85
		*size = 85
		fmt.Println("New size value is", *size)
	}
	fmt.Println("\n-----------------------------------")
	{
		b := 255
		// infer *int as type automatically
		var a = &b
		// Access value via b
		fmt.Printf("Value of b is %d\n", b)
		// Access value via dereferencing a
		fmt.Printf("Value of b is %d\n", *a)
		// ------------ Set value of b to another value ------------
		b = 69
		// Access value via b
		fmt.Printf("Value of b is %d\n", b)
		// Access value via dereferencing a
		fmt.Printf("Value of b is %d\n", *a)
	}
	fmt.Println("\n-----------------------------------")
	{
		// In GO we can return pointers from functions, as GO will allocate the memory for it on the heap.
		// In C or C++ this would not be possible
		pointer := getAnIntPointer()
		fmt.Printf("Points to %v and the value is %d\n", pointer, *pointer)
	}
	fmt.Println("\n-----------------------------------")
	{
		pointer := getAnIntPointer()
		pointerPointer := &pointer
		fmt.Printf("Pointer                                      points at %v   and the value is %d\n", pointer, *pointer)
		fmt.Printf("pointerPointer points at %v and the value is %v, which resolves to %d\n", pointerPointer, *pointerPointer, **pointerPointer)
		b := 6
		pointer = &b
		fmt.Printf("Pointer                                      points at %v   and the value is %d\n", pointer, *pointer)
		fmt.Printf("pointerPointer points at %v and the value is %v, which resolves to %d\n", pointerPointer, *pointerPointer, **pointerPointer)
	}
}

func getAnIntPointer() *int {
	i := 5
	return &i
}
