package rectangle

import (
	"math"
	"fmt"
)

func init() {
	fmt.Println("rectangle package initialized")
}

func Area(len, wid float64) float64 {
	area := len * wid
	return area
}

func Diagonal(len, wid float64) float64 {
	diagonal := math.Sqrt((len * len) + (wid * wid))
	return diagonal
}

// unexported function, because it starts with a lowerCase
func internalFunction() (value int) {
	value = 5
	return;
}
