package main

import (
	"fmt"
)

var LEDS = map[string]int{
	"1": 2,
	"2": 5,
	"3": 5,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 3,
	"8": 7,
	"9": 6,
	"0": 6,
}

func calculate(input string) int {
	out := 0
	for _, j := range input {
		out += LEDS[string(j)]
	}
	return out
}

func main() {
	var n int
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Println("input error in row quantity", err)
		return
	}

	for i := 0; i < n; i++ {
		var input string
		_, err := fmt.Scanf("%s", &input)
		if err != nil {
			fmt.Println("input error in rows", err)
			return
		}
		a := calculate(input)
		fmt.Printf("%d leds\n", a)
	}
}
