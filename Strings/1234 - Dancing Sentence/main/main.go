package main

import (
	"bufio"
	"os"
	"strings"
)

const numberOfLetters = 32

func dance(input string) string {
	inputBytes := []byte(strings.ToLower(input))
	result := make([]byte, len(inputBytes))

	uppercase := true
	for k, v := range inputBytes {
		if v == ' ' {
			result[k] = ' '
			continue
		}
		if uppercase {
			result[k] = v - numberOfLetters
			uppercase = false
			continue
		}
		result[k] = v
		uppercase = true
	}
	return string(result)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	writer := bufio.NewWriter(os.Stdout)

	for scanner.Scan() {
		input := scanner.Text()

		result := dance(input)
		writer.WriteString(result + "\n")
		writer.Flush()
	}
}
