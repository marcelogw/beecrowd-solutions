package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func hash(v string, max int) int {
	vv, _ := strconv.Atoi(v)
	return vv % max
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	cases, _ := strconv.Atoi(scanner.Text())

	for i := 1; i <= cases; i++ {
		scanner.Scan()
		params := strings.Fields(scanner.Text())
		m, _ := strconv.Atoi(params[0])

		table := make(map[int][]string, m)
		for j := 0; j < m; j++ {
			table[j] = []string{}
		}

		scanner.Scan()
		addresses := strings.Fields(scanner.Text())
		for _, v := range addresses {
			h := hash(v, m)
			table[h] = append(table[h], v)
		}

		for k := 0; k < m; k++ {
			v := table[k]
			fmt.Printf("%d -> ", k)
			if len(v) > 0 {
				fmt.Printf("%s -> ", strings.Join(v, " -> "))
			}
			fmt.Printf("\\\n")
		}

		if i < cases {
			fmt.Println()
		}
	}
}
