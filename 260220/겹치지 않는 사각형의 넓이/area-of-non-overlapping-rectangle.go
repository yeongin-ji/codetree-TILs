package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const size = 2100

func SetArea(a *[size][size]int, x1, y1, x2, y2, offset int) {
	for i := x1; i < x2; i++ {
		for j := y1; j < y2; j++ {
			a[i+1000][j+1000] += offset
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	nextInt := func() int {
		scanner.Scan()
		val, _ := strconv.Atoi(scanner.Text())
		return val
	}

	// Rectangle A
	ax1 := nextInt()
	ay1 := nextInt()
	ax2 := nextInt()
	ay2 := nextInt()

	// Rectangle B
	bx1 := nextInt()
	by1 := nextInt()
	bx2 := nextInt()
	by2 := nextInt()

	// Rectangle M
	mx1 := nextInt()
	my1 := nextInt()
	mx2 := nextInt()
	my2 := nextInt()

	// Please write your code here.
	area := [size][size]int{}
	SetArea(&area, ax1, ay1, ax2, ay2, 1)
	SetArea(&area, bx1, by1, bx2, by2, 1)
	SetArea(&area, mx1, my1, mx2, my2, -1)

	var rst int
	for i := range size {
		for j := range size {
			if area[i][j] == 1 {
				rst++
			}
		}
	}
	fmt.Print(rst)
}