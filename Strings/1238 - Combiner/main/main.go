package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func combine(a, b string) string {
	r := ""
	la := len(a)
	lb := len(b)

	// merge strings
	for i := 0; i < la; i++ {
		r += string(a[i])
		if i < lb {
			r += string(b[i])
		}
	}

	// in the case of second string greater than first
	for i := la; i < lb; i++ {
		r += string(b[i])
	}

	return r
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < n; i++ {
		scanner.Scan()
		input := scanner.Text()

		words := strings.Split(input, " ")
		writer.WriteString(combine(words[0], words[1]) + "\n")
	}
}
