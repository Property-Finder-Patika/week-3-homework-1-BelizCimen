package main

import "fmt"

func main() {
	// When comparing two arrays, their types should be identical

	// Comment out one of the following pair of variables and observe the results.

	var (
		// equal (types + elements are identical)::
		blue = [3]int{6, 9, 3}
		red  = [3]int{6, 9, 3}
	)
	fmt.Printf("blue bookcase : %v\n", blue)
	fmt.Printf("red bookcase  : %v\n", red)

	fmt.Println("Are they equal?", blue == red)
}
