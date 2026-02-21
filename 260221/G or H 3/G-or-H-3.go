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
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())
	pos := make([]int, n)
	letter := make([]string, n)
	maxVal := math.MinInt
	for i := 0; i < n; i++ {
		scanner.Scan()
		posVal, _ := strconv.Atoi(scanner.Text())
		pos[i] = posVal
		scanner.Scan()
		letter[i] = scanner.Text()
		maxVal = max(maxVal, posVal)
	}

	// Please write your code here.
	line := make([]string, maxVal+1)
	for i := range n {
		line[pos[i]] = letter[i]
	}
	maxVal = math.MinInt
	for i := 1; i < len(line)-k; i++ {
		sum := 0
		for j := range k+1 {
			sum += Score(line[i+j])
		}
		maxVal = max(maxVal, sum)
	}
	fmt.Print(maxVal)
}

func Score(s string) int {
	rst := 0
	if s == "G" {
		rst = 1
	} else if s == "H" {
		rst = 2
	}
	return rst
}