package main
import (
	"fmt"
	"tutorial/geometry/rectangle"
	"log"
	// package below is imported, but never used in this code. We just wanted the init method to do its magic.
	_ "tutorial/geometry/implicit"
)

var rectLen, rectWidth float64 = 6, 7

func init() {
	fmt.Println("geometry package initialized (init fun 1)")
	if rectLen < 0 {
		log.Fatal("length is less than zero")
	}
	if rectWidth < 0 {
		log.Fatal("width is less than zero")
	}
}

func init() {
	fmt.Println("geometry package initialized (init fun 2)")
}

func main() {
	var rectLen, rectWidth float64 = 6, 7
	fmt.Println("Geometrical shape properties")
	/*Area function of rectangle package used         */
	fmt.Printf("area of rectangle %.2f\n", rectangle.Area(rectLen, rectWidth))
	/*Diagonal function of rectangle package used         */
	fmt.Printf("diagonal of the rectangle %.2f ", rectangle.Diagonal(rectLen, rectWidth))
}

func init() {
	fmt.Println("geometry package initialized (init fun 3)")
}
