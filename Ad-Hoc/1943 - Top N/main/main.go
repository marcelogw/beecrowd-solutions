package main

import (
	"bufio"
	"os"
	"strconv"
)

var positions = [7]int{1, 3, 5, 10, 25, 50, 100}

func getTop(i int) int {
	for _, v := range positions {
		if i <= v {
			return v
		}
	}
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Scan()
	i, _ := strconv.Atoi(scanner.Text())
	_, _ = writer.WriteString("Top " + strconv.Itoa(getTop(i)) + "\n")
}
