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

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		arr[i], _ = strconv.Atoi(scanner.Text())
	}

	// Please write your code here.
	maxVal := math.MinInt
	for i := range n-k+1 {
		sum := 0
		for j := range k {
			sum += arr[i+j]
		}
		maxVal = max(maxVal, sum) 
	}
	fmt.Print(maxVal)
}