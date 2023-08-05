package main

import (
	"fmt"
)

// -- helpers --

func scanInt() int {
	var i int
	_, err := fmt.Scan(&i)
	if err != nil {
		panic(err)
	}
	return i
}

// -- pascal --

func pascal(n int) int {
	return 1 << (n - 1) // shift left to obtain the potency of two [2, 4, 8, 16...]
}

func sumPascal(n int) int {
	if n == 0 {
		return 1
	}
	return pascal(n) + sumPascal(n-1)
}

// -- main --

func main() {
	t := scanInt()
	for i := 0; i < t; i++ {
		sum := sumPascal(scanInt())
		fmt.Println(sum)
	}
}
