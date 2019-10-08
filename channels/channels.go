package main

import (
	"fmt"
	"time"
)

func hello(done chan bool) {
	fmt.Println("Going to sleep soon")
	// Emphasizes the blocking WAIT on our main ;)
	time.Sleep(2 * time.Second)
	fmt.Println("I love the smell of napalm in the morning")
	done <- true
}

func squares(result chan int, input ...int) {
	for _, v := range input {
		_result :=  v * v
		fmt.Printf("(%d * %d) = %d\n", v, v, _result)
		result <- _result
	}
}

func cubes(result chan int, input ...int) {
	for _, v := range input {
		_result :=  v * v * v
		fmt.Printf("(%d * %d * %d) = %d\n", v, v, v, _result)
		result <- _result
	}
}

func main() {
	{
		var a chan int
		if a == nil {
			fmt.Println("channel a is nil, going to define it")
			a = make(chan int)
			fmt.Printf("Type of a is %T\n", a)
		}
	}
	{
		b := make(chan bool)
		// Call the function as a "Go-Routine"
		go hello(b)
		// Read in from the channel "b", blocking - waits until we receive a value
		// Without that statement our MAIN will most probably terminate before our GO Routine had finished.
		<-b
		fmt.Println("done")
	}
	fmt.Println("\n-------------------------------")
	{
		result := make(chan int)
		input := []int{1, 2, 3, 4, 5, 6, 7}
		// Call routines
		go squares(result, input...)
		go cubes(result, input...)
		// Collect
		sum := 0
		for i := 0; i < len(input) * 2; i++ {
			sum += <- result
		}
		fmt.Println("Total:", sum)
	}
}
