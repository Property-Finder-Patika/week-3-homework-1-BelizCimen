package main

import (
	"fmt"
	"strings"
)

func main() {

	// defining integer array
	s := make([]int, 5)
	s[0] = 1
	s[1] = 2
	s[2] = 3
	s[3] = 4
	s[4] = 5
	fmt.Println(s)

	fmt.Println("enter rotating count ")
	var count int
	fmt.Scanln(&count)

	fmt.Println("rotate left or right? l for left r for right ")
	fmt.Println("l / r ? ")

	var direction string
	fmt.Scanln(&direction)

	if strings.ToLower(direction) == "l" {
		fmt.Println(leftRotate(s, count))
	} else if strings.ToLower(direction) == "r" {
		fmt.Println(rightRotate(s, count))
	} else {
		fmt.Errorf("please enter r for right and l for left")
	}

}

// function rotate based on the count
func leftRotate(s []int, count int) []int {
	first := s[count:]
	last := s[:count]
	var fin []int
	fin = append(first, last...)
	return fin
}

func rightRotate(s []int, count int) []int {

	rot := len(s) - count
	first := s[rot:]
	last := s[:rot]
	var fin []int
	fin = append(first, last...)
	return fin
}
