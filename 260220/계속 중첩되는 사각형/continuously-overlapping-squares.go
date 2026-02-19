package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const size = 210

func SetRect(a *[size][size]int, x1, y1, x2, y2, color int) {
	for i := x1; i < x2; i++ {
		for j := y1; j < y2; j++ {
			a[i+100][j+100] = color
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	x1 := make([]int, n)
	y1 := make([]int, n)
	x2 := make([]int, n)
	y2 := make([]int, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		x1[i], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		y1[i], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		x2[i], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		y2[i], _ = strconv.Atoi(scanner.Text())
	}

	// Please write your code here.
	area := [size][size]int{}
	for i := range n {
		color := i%2 + 1 // 1: red, 2: blue
		SetRect(&area, x1[i], y1[i], x2[i], y2[i], color)
	}
	cnt := 0
	for i := range size {
		for j := range size {
			if area[i][j] == 2 {
				cnt++
			}
		}
	}
	fmt.Print(cnt)
}