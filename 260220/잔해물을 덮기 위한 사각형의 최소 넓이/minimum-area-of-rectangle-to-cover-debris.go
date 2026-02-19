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

	scanner.Scan()
	r1x1, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	r1y1, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	r1x2, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	r1y2, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	r2x1, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	r2y1, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	r2x2, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	r2y2, _ := strconv.Atoi(scanner.Text())

	// Please write your code here.
	area := [size][size]int{}
	SetArea(&area, r1x1, r1y1, r1x2, r1y2, 1)
	SetArea(&area, r2x1, r2y1, r2x2, r2y2, -1)
	
	var minFlag bool
	var width, height int
	var minx, maxx, miny, maxy int
	for i := range size {
		for j := range size {
			if area[i][j] == 1 {
				if !minFlag {
					minx = i
					miny = j
					minFlag = true
				}
				if i > maxx {
					maxx = i
				}
				if j > maxy {
					maxy = j
				}
			}
		}
	}
	width = maxx - minx + 1
	height = maxy - miny + 1

	fmt.Print(width*height)
}