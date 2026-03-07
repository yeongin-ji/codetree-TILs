package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var dp []int

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	// Please write your code here.
	dp = make([]int, n+1)
	for i := 1; i <= n; i++ {
		if i <= 2 {
			dp[i] = 1
		} else {
			dp[i] = dp[i-1] + dp[i-2]
		}
	}
	fmt.Print(dp[n])
}