package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"slices"
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
	for i := range n {
		if i%2==0 {
			tmp := slices.Clone(arr[:i+1])
			slices.Sort(tmp)
			mid_i := len(tmp)/2
			fmt.Printf("%d ", tmp[mid_i])
		}
	}
}