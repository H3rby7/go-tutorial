package main

import (
	"fmt"
	"tutorial/structures/computer"
)

type Employee2 struct {
	firstName string
	lastName  string
	age       int
}

type Employee struct {
	firstName, lastName string
	age, salary         int
}

type FullEmployee struct {
	name string
	computer.Spec
}

func main() {
	{
		susan := Employee{
			firstName: "Susan",
			lastName:  "Becks",
			age:       4,
			salary:    20,
		}
		fmt.Println(susan)
		fmt.Printf("%s %s is %d years old and gets $%d of pocket money each month\n", susan.firstName, susan.lastName, susan.age, susan.salary)
	}
	fmt.Println("\n-----------------------------------")
	{
		susan := Employee{"Susan", "Becks", 4, 20}
		fmt.Println(susan)
	}
	fmt.Println("\n-----------------------------------")
	{
		employee := struct {
			firstName, lastName string
			age, salary         int
		}{"Peter", "Griffin", 52, 0}
		fmt.Printf("Employee %v using anonymous struct %T\n", employee, employee)
	}
	fmt.Println("\n-----------------------------------")
	{
		empl := Employee{}
		// Initialized with default values
		fmt.Printf("Nil-Employee %v\n", empl)
		var emp2 Employee
		fmt.Printf("Nil-Employee2 %v\n", emp2)
		emp2.firstName = "Not"
		emp2.lastName = "Sure"
		fmt.Printf("Nil-Employee2 with assigned values %v\n", emp2)
	}
	fmt.Println("\n-----------------------------------")
	{
		empl := Employee{salary: 3200, firstName: "Kra"}
		fmt.Printf("Employee %v\n", empl)
	}
	fmt.Println("\n-----------------------------------")
	{
		employeePointer := &Employee{salary: 3200, firstName: "Kra"}
		// Note how %v gives us some insight to the object we are referring to, while %p gives us the address.
		fmt.Printf("Pointer to %p (%v), where an employee resides %v\n", employeePointer, employeePointer, *employeePointer)
	}
	fmt.Println("\n-----------------------------------")
	{
		emp := FullEmployee{
			name: "Susan",
			Spec: computer.Spec{
				Maker: "Intel",
				Price: 1,
				// property below is not exported, so we cannot set it.
				// model: "heheh",
			},
		}
		fmt.Println(emp)
		// We can jump to the 'Maker' field
		fmt.Println("Maker is ", emp.Maker)
	}
	fmt.Println("\n-----------------------------------")
	fmt.Println("Two structs can be compared by '==' if all of their fields are comparable by '=='")
}
