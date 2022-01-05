package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var r1 = regexp.MustCompile(`\d+`)
var r2 = regexp.MustCompile(`\b[a-z]`)


func formatName(line string, year int) string {
	initials := r2.FindAllString(line, -1)

	return fmt.Sprintf("%s%d", strings.Join(initials, ""), year)
}

func getConflicts(r *bufio.Reader, n, year int) int {
	usersNames := map[string]bool{}
	conflicts := 0
	for i := 0; i < n; i++ {
		line, err := r.ReadString('\n')
		if err != nil {
			fmt.Printf("failed to read the line: %s", err)
			panic(err)
		}

		username := formatName(line, year)

		_, ok := usersNames[username]
		if !ok {
			usersNames[username] = true
		}else {
			conflicts++
		}
	}
	return conflicts
}

func getValues(line string) (int, int){
	values := r1.FindAllString(line, -1)
	n, err := strconv.Atoi(values[0])
	if err != nil {
		fmt.Printf("failed to parse first value to int: %s", err)
		panic(err)
	}
	year, err := strconv.Atoi(values[1])
	if err != nil {
		fmt.Printf("failed to parse second value to int: %s", err)
		panic(err)
	}
	return n, year
}

func main(){
	r := bufio.NewReader(os.Stdin)
	for true {
		line, err := r.ReadString('\n')
		if err != nil || line == "" || line == "\n" {
			return
		}

		n, year := getValues(line)

		conflicts := getConflicts(r, n, year)

		fmt.Println(conflicts)
	}
}
