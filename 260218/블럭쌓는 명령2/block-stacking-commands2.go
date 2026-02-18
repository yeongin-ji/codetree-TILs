package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"slices"
)

func Simul(area []int, a, b int) {
	for i := a; i <= b; i++ {
		area[i]++
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())

	a := make([]int, k)
	b := make([]int, k)
	for i := 0; i < k; i++ {
		scanner.Scan()
		a[i], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		b[i], _ = strconv.Atoi(scanner.Text())
	}

	// Please write your code here.
	area := make([]int, n+10)
	for i := range k {
		Simul(area, a[i], b[i])
	}

	fmt.Print(slices.Max(area))
}