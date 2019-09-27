package main  

import "fmt"
import "math"

func rectProps(length, width float64)(float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

func rectProps2(length, width float64)(area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return
}


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
	const price = 10
	fmt.Println("Price for 5 items is", calculateBill(price, count))
	fmt.Println("Price for 5 items is", calculateBill2(price, count))
	var area, perimeter = rectProps(float64(width), float64(height))
	var areaOnly, _ = rectProps2(float64(width), float64(height))
	fmt.Println("a", width, "x", height, "rectangle has an area of", area, "and perimeter", perimeter)
	fmt.Println("a", width, "x", height, "rectangle has an area of", areaOnly)
}

func calculateBill(price int, no int) int {
	var totalPrice = price * no
	return totalPrice
}

func calculateBill2(price, no int) int {
	return price * no
}
