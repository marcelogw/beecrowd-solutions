package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"
)

func calcMinutes(ch, cm, ah, am int) int {
	current := time.Date(0, 0, 0, ch, cm, 0, 0, time.UTC)
	alarm := time.Date(0, 0, 0, ah, am, 0, 0, time.UTC)

	sub := alarm.Sub(current).Minutes()
	if sub <= 0 {
		alarm = alarm.Add(24 * time.Hour)
		sub = alarm.Sub(current).Minutes()
	}
	return int(sub)
}

func toInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for scanner.Scan() {
		times := strings.Fields(scanner.Text())
		ch, cm, ah, am := toInt(times[0]), toInt(times[1]), toInt(times[2]), toInt(times[3])
		if ch == 0 && cm == 0 && ah == 0 && am == 0 {
			break
		}

		result := calcMinutes(ch, cm, ah, am)
		writer.WriteString(strconv.Itoa(result) + "\n")
	}
}
