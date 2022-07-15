package main

import "fmt"

func sieveOfEra(num int) []int {

	if num <= 2 {
		fmt.Println("not found")
	}

	sieve := make([]bool, num)
	primes := make([]int, 0)
	sieve[0], sieve[1] = true, true
	for k, sv := range sieve {
		if sv {
			continue
		}
		primes = append(primes, k)
		for n := 2; n*k < num; n++ {
			sieve[n*k] = true
		}
	}
	return primes
}

func main() {
	fmt.Println("Enter the number ")
	var num int
	fmt.Scanln(&num)
	fmt.Println(sieveOfEra(num))
}
