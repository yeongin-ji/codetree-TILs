package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func SetArea(a *[210][210]int, x1, y1, x2, y2 int) {
	for i := x1; i < x2; i++ {
		for j := y1; j < y2; j++ {
			a[i+100][j+100]++
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	nextInt := func() int {
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		return v
	}

	n := nextInt()
	x1 := make([]int, n)
	y1 := make([]int, n)
	x2 := make([]int, n)
	y2 := make([]int, n)
	for i := 0; i < n; i++ {
		x1[i] = nextInt()
		y1[i] = nextInt()
		x2[i] = nextInt()
		y2[i] = nextInt()
	}

	// Please write your code here.
	area := [210][210]int{}
	for i := range n {
		SetArea(&area, x1[i], y1[i], x2[i], y2[i])
	}
	var rst int
	for i := range 210 {
		for j := range 210 {
			if area[i][j] > 0 {
				rst++
			}
		}
	}
	fmt.Print(rst)
}