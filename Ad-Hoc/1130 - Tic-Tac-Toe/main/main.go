package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func checkWin(line string) string {
	// irregulars cases
	if strings.Count(line, "XX") > 0 || strings.Contains(line, "X.X") {
		return "S"
	}

	//_, line, _ := strings.Cut(line, "X")
	//isEven := (l % 2) == 0
	//if isEven {
	//	i := strings.Index(line, "X")
	//	if i < 0 {
	//		return "S"
	//	}
	//}

	// remove plays
	line = strings.ReplaceAll(line, "..X..", "")
	line = strings.ReplaceAll(line, ".X..", "")
	line = strings.ReplaceAll(line, "..X.", "")
	line = strings.ReplaceAll(line, "..X", "")
	line = strings.ReplaceAll(line, "X..", "")
	l := len(line)

	if l == 0 {
		return "N"
	}

	// empty table
	div := l / 5
	isDivEven := div%2 == 0
	remainder := l % 5
	//if {
	if (isDivEven && remainder > 0) || (!isDivEven && remainder == 0) {
		return "S"
	} else {
		return "N"
	}
	//}

	//return "N"
}

func main() {
	//reader := bufio.NewReader(os.Stdin)
	reader, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(reader)
	output, _ := os.Create("output.txt")
	writer := bufio.NewWriter(output)
	defer writer.Flush()

	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		if n == 0 {
			return
		}
		scanner.Scan()
		line := scanner.Text()
		result := checkWin(line)
		//fmt.Println(result)
		writer.WriteString(result + "\n")
	}

}
