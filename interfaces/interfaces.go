package main

import (
	"fmt"
)

//interface definition
type VowelsFinder interface {
	FindVowels() []rune
}

type MyString string

//MyString implements VowelsFinder
func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
		if rune == 'A' || rune == 'E' || rune == 'I' || rune == 'O' || rune == 'U' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

// This interface has no methods, is thus implemented by all things.
func describe(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

func assertInt(i interface{}) {
	// We can assert the type and catch an error by using the 'ok' or let's just say a second variable upon asserting
	value, ok := i.(int)
	if ok {
		fmt.Println("It is an int!", value)
		return
	}
	// Value will always have the default value here - thus 0
	fmt.Println("Not an int!", value)
}

func findType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("It's an INT")
	case float64, float32:
		fmt.Println("It's a float")
	// We can also check for the interface!
	case VowelsFinder:
		fmt.Println("It can find vowels")
	default:
		fmt.Println("Well I don't know what it is")
	}
}

func main() {
	{
		name := MyString("Sam Anderson")
		var v VowelsFinder
		v = name // possible since MyString implements VowelsFinder
		fmt.Printf("Vowels are %c", v.FindVowels())
	}
	fmt.Println("\n-----------------------------------")
	{
		// All these vars and structs implement the empty interface and can thus use the describe function
		s := "Hello World"
		describe(s)
		i := 55
		describe(i)
		strt := struct {
			name string
		}{
			name: "Naveen R",
		}
		describe(strt)
	}
	fmt.Println("\n-----------------------------------")
	{
		var s interface{} = "Steven Paul"
		assertInt(s)
		var a interface{} = 69
		assertInt(a)
	}
	fmt.Println("\n-----------------------------------")
	{
		findType(5)
		findType(42.0)
		findType(MyString("im a string. yeah!"))
		findType("im a string. yeah!")
	}
}
