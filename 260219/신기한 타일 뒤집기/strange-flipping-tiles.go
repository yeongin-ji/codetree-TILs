package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func SetTile(t []int, x int, dir string, pos int) int {
	var idx, sign int
	if dir == "L" {
		sign = -1
	} else if dir == "R" {
		sign = 1
	}
	for i := range x {
		idx = pos + i*sign + 100_000
		t[idx] = sign
	}
	return idx - 100_000
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	xs := make([]int, n)
	dirs := make([]string, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		xs[i], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		dirs[i] = scanner.Text()
	}

	// Please write your code here.
	tiles := make([]int, 200_200)
	var idx int
	for i := range n {
		idx = SetTile(tiles, xs[i], dirs[i], idx)
	}
	var wt, bl int
	for _, v := range tiles {
		if v == -1 {
			wt++
		} else if v == 1 {
			bl++
		}
	}
	fmt.Print(wt, bl)
}