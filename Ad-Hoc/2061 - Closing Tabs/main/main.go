package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	params := strings.Fields(scanner.Text())
	tabs, _ := strconv.Atoi(params[0])
	actions, _ := strconv.Atoi(params[1])
	for i := 0; i < actions; i++ {
		scanner.Scan()
		action := scanner.Text()
		if action == "fechou" {
			tabs++
			continue
		}
		tabs--
	}
	fmt.Println(tabs)
}
