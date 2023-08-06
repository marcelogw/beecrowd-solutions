package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func toInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func getResult(m map[int]bool) string {
	sum := 0
	for _, v := range m {
		if v {
			sum++
		}
	}
	return strconv.Itoa(sum)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for scanner.Scan() {
		params := strings.Fields(scanner.Text())
		contestants, problems := toInt(params[0]), toInt(params[1])
		if contestants == 0 && problems == 0 {
			break
		}

		passing := map[int]bool{
			0: true,
			1: true,
			2: true,
			3: true,
		}
		solutions := make(map[int]int, problems)
		for i := 1; i <= contestants; i++ {
			scanner.Scan()
			final := scanner.Text()
			if passing[0] && !strings.Contains(final, "0") {
				passing[0] = false
			}
			if passing[1] && !strings.Contains(final, "1") {
				passing[1] = false
			}

			f := strings.Fields(final)
			for k, v := range f {
				n, _ := strconv.Atoi(v)
				solutions[k] += n
			}
		}

		for _, v := range solutions {
			if passing[2] && v == 0 {
				passing[2] = false
			}
			if passing[3] && v == contestants {
				passing[3] = false
			}
			if !passing[2] && !passing[3] {
				break
			}
		}

		writer.WriteString(getResult(passing) + "\n")
	}
}
