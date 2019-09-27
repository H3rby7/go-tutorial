package main  

import "fmt"
import "math"
func main() {
	var age int
	fmt.Println("My age is", age)
	age = 10
	fmt.Println("My age is", age)
	age = 20
	fmt.Println("My age is", age)
	var year = 2020
	fmt.Println("The year is", year)
	var width, height = 20, 10
	fmt.Println("Width is", width, "Height is", height)
	var (
		color = "red"
		count = 4
		id int
	)
	fmt.Println("The color is", color, "of", count, "pieces with id", id)
	name, digit := "peter", 7
	fmt.Println(name, "loves the number", digit)
	age1, age2 := 23.0, 27
	lowestAge := math.Min(age1, float64(age2))
	fmt.Println("The youngest person is", lowestAge, "old")

	const typedhello string = "Hello World"
	fmt.Println(typedhello)
	var defaultName = "Sam" //allowed         
	type myString string
	// assigning 'Sammy' to the variable customName of type myString
	var customName myString = "Sammy" //allowed
	//customName = defaultName     // not allowed, because customName is of type myString, not of type string
	fmt.Println("defaultName:", defaultName, "customName:", customName)
}

