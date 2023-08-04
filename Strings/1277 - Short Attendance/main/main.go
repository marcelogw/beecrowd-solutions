package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkAttendance(students, grades []string) []string {
	var result []string

	for index, grade := range grades {
		attendances := 0
		faults := 0
		for _, g := range grade {
			if g == 'P' {
				attendances++
			}
			if g == 'A' {
				faults++
			}
		}
		percentage := (attendances * 100) / (faults + attendances)
		if percentage < 75 {
			result = append(result, students[index])
		}
	}

	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < n; i++ {
		scanner.Scan() // we can ignore number of students
		scanner.Scan()
		students := strings.Split(scanner.Text(), " ")

		scanner.Scan()
		grades := strings.Split(scanner.Text(), " ")

		result := checkAttendance(students, grades)
		fmt.Println(strings.Join(result, " "))
	}
}
