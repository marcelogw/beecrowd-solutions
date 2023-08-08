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

type Queue struct {
	q     []int
	index int
}

func (q *Queue) Add(i int) {
	q.index++
	q.q[q.index] = i
}

func (q *Queue) Pop() int {
	if q.index > 0 {
		v := q.q[q.index]
		q.index--
		return v
	}
	return 0
}

func (q *Queue) Current() int {
	return q.q[q.index]
}

func newQueue(i int) Queue {
	return Queue{
		q: make([]int, i),
	}
}

func checkPos(f []int) bool {
	q := newQueue(len(f))
	n := len(f)
	for i := n - 1; i >= 0; {
		if q.Current() == n {
			q.Pop()
			n--
			continue
		}
		if f[i] == n {
			n--
			i--
			continue
		}
		q.Add(f[i])
		i--
	}
	for i := q.index; i > 0; i-- {
		if q.Pop() == n {
			n--
			continue
		}
		return false
	}
	return n == 0
}

func convert(s []string) []int {
	res := make([]int, len(s))
	for k, v := range s {
		res[k] = toInt(v)
	}
	return res
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for scanner.Scan() {
		n := toInt(scanner.Text())
		if n == 0 {
			break
		}
		for scanner.Scan() {
			f := strings.Fields(scanner.Text())
			if len(f) == 1 && f[0] == "0" {
				break
			}
			res := checkPos(convert(f))
			if res {
				writer.WriteString("Yes\n")
				continue
			}
			writer.WriteString("No\n")
		}
		writer.WriteString("\n")
	}
}
