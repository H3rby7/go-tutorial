package main

import (
	"fmt"
	"unicode/utf8"
)

func printBytes(s string) {
	for i := 0; i < len(s); i++ {
		// %x to print Hexa-Decimals
		fmt.Printf("%x ", s[i])
	}
}

func printChars(s string) {
	for i := 0; i < len(s); i++ {
		// %c to print CHARS
		fmt.Printf("%c ", s[i])
	}
}

func printRune(s string) {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
}

func printCharsAndBytes(s string) {
	// Range on the string works well to show runes and their respective start-byte in the string
	for bytePosition, _rune := range s {
		fmt.Printf("%c starts at byte %d\n", _rune, bytePosition)
	}
}

func main() {
	{
		name := "Hello World"
		printBytes(name)
		fmt.Printf("\n")
		printChars(name)
		fmt.Println("\n---------------------------")
		// note the ñ character uses 2 bytes in UTF-8, so the CHAR printing will look funny, because it assumes 1 byte per char.
		name = "Señor"
		printBytes(name)
		fmt.Printf("\n")
		printChars(name)
		fmt.Println("\n---------------------------")
		// PrintRune represents a UniCode Point and handles multiple bytes per character.
		printRune(name)
		fmt.Println("\n---------------------------")
		printCharsAndBytes(name)
	}
	fmt.Println("\n---------------------------")
	{
		byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
		// Construct a string from bytes
		str := string(byteSlice)
		fmt.Println(str)
	}
	fmt.Println("\n---------------------------")
	{
		byteSlice := []byte{67, 97, 102, 195, 169}
		// Construct a string from bytes
		str := string(byteSlice)
		fmt.Println(str)
	}
	fmt.Println("\n---------------------------")
	{
		runeSlice := []rune{0x43, 0x61, 0x66, 0xE9}
		// Construct a string from bytes
		str := string(runeSlice)
		fmt.Println(str)
	}
	fmt.Println("\n---------------------------")
	{
		word1 := "Señor"
		printLength(word1)
		word2 := "Pets"
		printLength(word2)
	}
	fmt.Println("\n---------------------------")
	{
		h := "hello"
		// Because strings are immutable we need to convert them to a slice before we 'edit' them
		fmt.Println(mutate([]rune(h)))
	}
}

func printLength(s string) {
	fmt.Printf("length (in runes)  of %s is %d\n", s, utf8.RuneCountInString(s))
	fmt.Printf("length (using len) of %s is %d\n", s, len(s))
}

func mutate(s []rune) string {
	s[0] = 'B'
	return string(s)
}