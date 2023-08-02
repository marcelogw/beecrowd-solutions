package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var lowerLimit = int('A')
var nLetters = 26

func decode(input string, shift int) string {
	res := ""
	for _, v := range input {
		s := int(v) - shift
		if s < lowerLimit {
			s += nLetters
		}
		res += string(rune(s))
	}
	return res
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)

	scanner.Scan()
	t := scanner.Text()
	n, err := strconv.Atoi(t)
	if err != nil {
		panic(err)
	}
	for i := 1; i <= n; i++ {
		scanner.Scan()
		input := scanner.Text()

		scanner.Scan()
		shift, _ := strconv.Atoi(scanner.Text())

		fmt.Println(decode(input, shift))
	}
}
