package main

import "fmt"

func main() {
	{
		// its length is part of its type
		var nums [5]int
		fmt.Printf("nums array: %#v\n", nums)
	}

	{
		// its length is not part of its type
		var nums []int
		fmt.Printf("nums slice: %#v\n", nums)

		fmt.Printf("len(nums) : %d\n", len(nums))

	}
}
