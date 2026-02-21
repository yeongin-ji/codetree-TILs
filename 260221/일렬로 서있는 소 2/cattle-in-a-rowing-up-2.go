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
	n, _ := strconv.Atoi(scanner.Text())

	A := make([]int, n+1)
	for i := 1; i <= n; i++ {
		scanner.Scan()
		A[i], _ = strconv.Atoi(scanner.Text())
	}

	// Please write your code here.
	cnt := 0
	for i := 1; i <= n; i++ {
		for j := i+1; j <= n; j++ {
			for k := j+1; k <= n; k++ {
				if A[i] <= A[j] && A[j] <= A[k] {
					cnt++
				}
			}
		}
	}
	fmt.Print(cnt)
}