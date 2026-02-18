package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"slices"
)

func SetLine(line []int, a, b int) {
	for i := a; i < b; i++ {
		line[i+100]++
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	x1 := make([]int, n)
	x2 := make([]int, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		x1[i], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		x2[i], _ = strconv.Atoi(scanner.Text())
	}

	// Please write your code here.
	line := make([]int, 210)
	for i := range n {
		SetLine(line, x1[i], x2[i])
	}
	fmt.Print(slices.Max(line))
}