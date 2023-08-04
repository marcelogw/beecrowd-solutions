package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

var OnlyWordsRegex = regexp.MustCompile(`\b[a-zA-Z]+\b`)

var filters = []*regexp.Regexp{
	regexp.MustCompile(`[a-zA-Z]+\.[a-zA-Z]+`), // remove words with dots between
	regexp.MustCompile(`\.[a-zA-Z]+`),          // remove words with dots before
	regexp.MustCompile(`[a-zA-Z]+\.{2,}`),      // remove words with more than 2 dots after
}

func filter(s string) string {
	for _, v := range filters {
		s = v.ReplaceAllString(s, "")
	}
	return s
}

func getDifficulty(input string) int {
	input = filter(input)
	matches := OnlyWordsRegex.FindAllString(input, -1)

	if len(matches) == 0 {
		return 250
	}

	sum := 0
	for _, v := range matches {
		sum += len(v)
	}
	avg := sum / len(matches)

	if avg <= 3 {
		return 250
	}
	if avg == 4 || avg == 5 {
		return 500
	}
	return 1000
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	writer := bufio.NewWriter(os.Stdout)

	for scanner.Scan() {
		input := scanner.Text()

		writer.WriteString(strconv.Itoa(getDifficulty(input)) + "\n")
		writer.Flush()
	}
}
