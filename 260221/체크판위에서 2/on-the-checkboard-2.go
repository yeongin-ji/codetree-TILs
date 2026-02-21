package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	R, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	C, _ := strconv.Atoi(scanner.Text())

	grid := make([][]string, R)
	for i := 0; i < R; i++ {
		grid[i] = make([]string, C)
		for j := 0; j < C; j++ {
			scanner.Scan()
			grid[i][j] = scanner.Text()
		}
	}

	// Please write your code here.
	start := grid[0][0]
	cnt := 0
	for i := 1; i < R; i++ {
		for j := 1; j < C; j++ {
			if mid1 := Reverse(start); grid[i][j] == mid1 {
				for k := i+1; k < R; k++ {
					for l := j+1; l < C; l++ {
						if mid2 := Reverse(mid1); grid[k][l] == mid2 {
							condition := k < R-1 && l < C-1 && grid[R-1][C-1] == Reverse(mid2)
							if condition {
								cnt++
							}
						}
					}
				}
			}
		}
	}
	fmt.Print(cnt)
}

func Reverse(prev string) string {
	rst := "W"
	if prev == "W" {
		rst = "B"
	}
	return rst
}