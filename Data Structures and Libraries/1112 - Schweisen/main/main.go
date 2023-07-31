package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"sync"
)

// --- helpers ----

func toInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// --- BIT ---

type BIT struct {
	tree  [][]int
	sizeN int
	sizeM int
	rw    *sync.RWMutex
}

func createBIT(sizeN, sizeM int) BIT {
	b := BIT{
		tree:  make([][]int, sizeN+1),
		sizeN: sizeN,
		sizeM: sizeM,
		rw:    &sync.RWMutex{},
	}
	for i := 0; i <= sizeN; i++ {
		b.tree[i] = make([]int, sizeM+1)
	}
	return b
}

func (b *BIT) sum(x, y int) int {
	sum := 0
	for i := x; i >= 0; i = (i & (i + 1)) - 1 {
		for j := y; j >= 0; j = (j & (j + 1)) - 1 {
			b.rw.RLock()
			sum += b.tree[i][j]
			b.rw.RUnlock()
		}
	}
	return sum
}

func (b *BIT) Query(x1, y1, x2, y2 int) int {
	x1_ := min(x1, x2)
	x2_ := max(x1, x2)
	y1_ := min(y1, y2)
	y2_ := max(y1, y2)

	return b.sum(x2_, y2_) - b.sum(x1_-1, y2_) - b.sum(x2_, y1_-1) + b.sum(x1_-1, y1_-1)
}

func (b *BIT) Update(x, y, value int) {
	for i := x; i <= b.sizeN; i = i | (i + 1) {
		for j := y; j <= b.sizeM; j = j | (j + 1) {
			b.rw.Lock()
			b.tree[i][j] += value
			b.rw.Unlock()
		}
	}
	return
}

// --- main ---

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for scanner.Scan() {
		params := strings.Fields(scanner.Text())
		sizeN := toInt(params[0])
		sizeM := toInt(params[1])
		price := toInt(params[2])

		if sizeN == 0 && sizeM == 0 && price == 0 {
			break
		}
		b := createBIT(sizeN, sizeM)

		scanner.Scan()
		n := toInt(scanner.Text())
		wg := sync.WaitGroup{}
		for i := 0; i < n; i++ {
			scanner.Scan()
			action := strings.Fields(scanner.Text())

			form := action[0]
			if form == "A" {
				n := toInt(action[1]) * price
				x := toInt(action[2]) + 1
				y := toInt(action[3]) + 1

				wg.Add(1)
				go func(x, y, n int) {
					b.Update(x, y, n)
					wg.Done()
				}(x, y, n)

			} else {
				x1 := toInt(action[1]) + 1
				y1 := toInt(action[2]) + 1
				x2 := toInt(action[3]) + 1
				y2 := toInt(action[4]) + 1

				wg.Wait()
				sum := b.Query(x1, y1, x2, y2)
				writer.WriteString(strconv.Itoa(sum) + "\n")
			}
		}
		writer.WriteString("\n")
	}
}
