package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"math"
)

func Abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		scanner.Scan()
		a[i], _ = strconv.Atoi(scanner.Text())
	}

	// 각 집이 선택되었다고 하나씩 가정
	m := math.MaxInt
	for i := 1; i <= n; i++ {
		sum := 0
		for j := 1; j <= n; j++ {
			sum += Abs(i-j) * a[j]
		}
		m = min(sum, m)
	}
	fmt.Print(m)
}