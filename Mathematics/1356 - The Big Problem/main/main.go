package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func getMDC(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func isSequence(a, b, c int64) bool {
	return 2*b == a+c
}

func sumSequence(terms int64, first int64, den int64) int64 {
	p1 := float64(terms) / 2
	p2 := float64(2*first + (terms-1)*den)
	return int64(p1 * p2)
}

// returns index of first sequence, denominator and current sum
func findSequence(n int64) (int64, int64, int64) {
	sum := int64(0)
	sqrt := int64(math.Sqrt(float64(n)))
	for i := int64(1); i < n; i++ {
		a := n % i
		if i > sqrt {
			b := n % (i + 1)
			c := n % (i + 2)
			if isSequence(a, b, c) {
				return i, b - a, sum
			}
		}
		sum += a
	}
	return 0, 0, sum
}

func sumAllSequences(n, index, den int64) int64 {
	sum := int64(0)
	for index < n {
		r := n % index
		div := (r / -den) + 1

		sum += sumSequence(div, r, den)
		index += div
		den += 1
	}
	return sum
}

func calc(n int64) (int64, int64) {
	index, den, sum := findSequence(n)
	if index > 0 {
		sum += sumAllSequences(n, index, den)
	}

	totalCases := n * n
	mdc := getMDC(sum, totalCases)

	return sum / mdc, totalCases / mdc
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		n, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		cases, total := calc(int64(n))
		fmt.Printf("%d/%d\n", cases, total)
	}
}
