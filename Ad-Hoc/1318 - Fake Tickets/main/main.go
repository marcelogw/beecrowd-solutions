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

func checkDuplicates(nTickets int, tickets []string) string {
	a := make([]int, nTickets)
	// a[0] - no ticket found yet
	// a[1] - found a duplicated
	// a[2] - already summed
	sum := 0
	for _, v := range tickets {
		i := toInt(v) - 1
		if a[i] == 0 {
			a[i] = 1
			continue
		}
		if a[i] == 1 {
			sum++
			a[i] = 2
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
		s := strings.Fields(scanner.Text())
		n, m := toInt(s[0]), toInt(s[1])
		if n == 0 && m == 0 {
			break
		}
		scanner.Scan()
		tickets := strings.Fields(scanner.Text())
		writer.WriteString(checkDuplicates(n, tickets) + "\n")
	}
}
