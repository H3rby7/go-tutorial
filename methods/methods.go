package main

import "fmt"

type Employee struct {
	name     string
	salary   int
	currency string
}

/*
 displaySalary() method has Employee as the receiver type
*/
func (e Employee) displaySalary() {
	fmt.Printf("Salary of %s is %s%d\n", e.name, e.currency, e.salary)
}

// Using a pointer here to make changes visible to the caller.
func (e *Employee) changeName(name string) {
	e.name = name
}

func (e Employee) changeSalary(salary int) {
	e.salary = salary
}

// Fancy shit: we can also create methods on non-struct  types
// They have to be in the same package though, this is why we cannot enhance 'int' directly.
type myInt int

func (a myInt) add(b myInt) myInt {
	return a + b
}

func main() {
	{
		emp := Employee{"Maria", 15000, "$"}
		emp.displaySalary()
	}
	fmt.Println("\n-----------------------------------")
	{
		emp := Employee{"Maria", 15000, "$"}
		(&emp).changeName("Martha")
		emp.displaySalary()
		// The language knows we use a pointer to call the function, so we can omit the (&struct) syntax
		emp.changeName("Shari")
		emp.displaySalary()
	}
	fmt.Println("\n-----------------------------------")
	{
		emp := Employee{"Maria", 15000, "$"}
		// Note that we can call the changeSalary Method with either: pointer or value!
		// The reason: go will transfer this line into (*(&emp)).changeSalary
		(&emp).changeSalary(17000)
		emp.displaySalary()
		emp.changeSalary(16000)
		emp.displaySalary()
		// For this example it doesn't really help us, because the changes are not propagated to us
	}
	fmt.Println("\n-----------------------------------")
	{
		// Using the 'add' method of the type alias 'myInt'
		a, b := myInt(5), myInt(7)
		sum := a.add(b)
		fmt.Println("Sum is", sum)
	}
}
