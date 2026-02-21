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
	
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		arr[i], _ = strconv.Atoi(scanner.Text())
	}
	
	// Please write your code here.
	var cnt int
	for i := range n {
		for j := i; j < n; j++ {
			var sum int
			var avg float64
			for k := j; k < n; k++ {
				sum += arr[k]
			}
			avg = float64(sum) / float64(n-j)
			for k := j; k < n; k++ {
				if avg == float64(arr[k]) {
					cnt++
					break
				}
			}
		}
	}
	fmt.Print(cnt)
}