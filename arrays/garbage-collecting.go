package main

import "fmt"

/*
As long as an array is referenced by a slice, it will stay in memory.
If we have a large array and only use a small proportion of it using a slice directly, the large array is still kept in memory
To work around that we can use copy and create an array of only the needed size, so the large one can be garbage collected.
 */

func countries() []string {
	// Large Array
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	// Slice of the actually needed parts - thus far not more memory efficient than before
	neededCountries := countries[:len(countries)-2]
	// Create a new empty slice with an underlying array of the size required (len(neededCountries))
	countriesCpy := make([]string, len(neededCountries))
	// Copy the required countries into the new array
	copy(countriesCpy, neededCountries)
	// return the small copy.
	return countriesCpy
	// Afterwards the large array can be garbage collected
}
func main() {
	countriesNeeded := countries()
	fmt.Println(countriesNeeded)
}
