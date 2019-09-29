package main

import (
	"fmt"
)

func main() {
	{
		personSalary := make(map[string]int)
		personSalary["steve"] = 12000
		personSalary["Maria"] = 15000
		personSalary["mike"] = 9000
		fmt.Println("personSalary map contents:", personSalary)
	}
	fmt.Println("\n--------------------------------------")
	{
		personSalary := map[string]int{
			"steve": 12000,
			"Maria": 15000,
		}
		personSalary["mike"] = 9000
		fmt.Println("personSalary map contents:", personSalary)
		// Joe is not in the list, thus the INT default is used, which is 0
		fmt.Println("Salary of joe is", personSalary["joe"])
		salary, ok := personSalary["joe"]
		fmt.Println("Joe is present?", ok, "and has a salary of", salary)
		_, present := personSalary["holger"]
		fmt.Println("Holger is present?", present)
		_, present = personSalary["Maria"]
		fmt.Println("Maria is present?", present)
	}
	fmt.Println("\n--------------------------------------")
	{
		personSalary := map[string]int{
			"steve": 12000,
			"Maria": 15000,
			"Mike":  9000,
		}
		printMyMap(personSalary)
		fmt.Println("Let's fire Steve")
		delete(personSalary, "steve")
		printMyMap(personSalary)
		fmt.Println("Let's fire Konrad - just in case we have one")
		delete(personSalary, "Konrad")
		printMyMap(personSalary)
	}
	fmt.Println("\n--------------------------------------")
	{
		personSalary := map[string]int{
			"steve": 12000,
			"Maria": 15000,
			"Mike":  9000,
		}
		fmt.Println("Size of our map is", len(personSalary))
		newPersonSalary := personSalary
		printMyMap(personSalary)
		// Maps are referenceTypes, thus changes to the newMap reflect in the old map
		fmt.Println("Let's fire Mike")
		delete(newPersonSalary, "Mike")
		fmt.Println("Original MAP:")
		printMyMap(personSalary)
		fmt.Println("New MAP")
		printMyMap(newPersonSalary)
	}
	fmt.Println("\n--------------------------------------")
	compareSomeMapsTest()
}

func printMyMap(personSalary map[string]int) {
	for key, value := range personSalary {
		fmt.Printf("personSalary[%s] = %d\n", key, value)
	}
}

func compareMaps(mapA map[string]int, mapB map[string]int) bool {
	if mapA == nil && mapB == nil {
		return true
	}
	if (mapA == nil && mapB != nil) || (mapB == nil && mapA != nil) {
		return false
	}
	if lengthA, lengthB := len(mapA), len(mapB); lengthA != lengthB {
		return false
	}
	for keyA, valueA := range mapA {
		valueB, present := mapB[keyA]
		if !present {
			return false
		}
		if valueA != valueB {
			return false
		}
	}
	return true
}

func compareSomeMapsTest() {
}