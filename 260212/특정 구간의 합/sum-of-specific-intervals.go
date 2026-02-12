package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var A []int

func rangeSum(a1, a2 int) int {
	sum := 0
	for i := a1; i <= a2; i++ {
		sum += A[i]
	}
	return sum
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	arr := make([]int, n+1)
	for i := 1; i <= n; i++ {
		scanner.Scan()
		arr[i], _ = strconv.Atoi(scanner.Text())
	}

	queries := make([][2]int, m)
	for i := 0; i < m; i++ {
		scanner.Scan()
		a1, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		a2, _ := strconv.Atoi(scanner.Text())
		queries[i] = [2]int{a1, a2}
	}

	// Please write your code here.
	A = arr
	for _, q := range queries {
		fmt.Println(rangeSum(q[0], q[1]))
	}
}