package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func SetLine(line []int, dist int, dir string, pos int) int {
	var sign, idx, offset int
	if dir == "R" {
		sign = 1
		offset = 0
	} else if dir == "L" {
		sign = -1
		offset = 1
	}

	for i := range dist {
		idx = (pos + i*sign - offset) + 1000
		line[idx]++
		idx += 1-offset
	}
	return idx - 1000
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	distances := make([]int, n)
	directions := make([]string, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		d, _ := strconv.Atoi(scanner.Text())
		distances[i] = d
		scanner.Scan()
		directions[i] = scanner.Text()
	}

	// Please write your code here.
	line := make([]int, 2100)
	var idx int
	for i := range n {
		idx = SetLine(line, distances[i], directions[i], idx)
	}
	var cnt int
	for _, v := range line {
		if v >= 2 {
			cnt++
		}
	}
	fmt.Print(cnt)
}