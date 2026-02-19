package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const size = 210

func SetArea(a *[size][size]int, x, y int) {
	for i := x; i < 8; i++ {
		for j := y; j < 8; j++ {
			a[i+100][j+100]++
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
	
	for i := 0; i < n; i++ {
		scanner.Scan()
		x1[i], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		y1[i], _ = strconv.Atoi(scanner.Text())
	}
	
	// Please write your code here.
	area := [size][size]int{}
	for i := range n {
		SetArea(&area, x1[i], y1[i])
	}
	cnt := 0
	for i := range size {
		for j := range size {
			if area[i][j] > 0 {
				cnt++
			}
		}
	}
	fmt.Print(cnt)
}