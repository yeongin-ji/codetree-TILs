package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"math"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	// Read N
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	// Read grid values
	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, n)
		for j := 0; j < n; j++ {
			scanner.Scan()
			grid[i][j], _ = strconv.Atoi(scanner.Text())
		}
	}

	// Please write your code here.
	maxNum := math.MinInt
	for i := range n {
		for j := range n-2 {
			cnt := grid[i][j] + grid[i][j+1] + grid[i][j+2]
			maxNum = max(maxNum, cnt)
		}
	}
	fmt.Print(maxNum)
}