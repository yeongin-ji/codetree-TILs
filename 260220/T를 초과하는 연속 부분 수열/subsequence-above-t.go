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
	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		arr[i], _ = strconv.Atoi(scanner.Text())
	}

	// Please write your code here.
	list := make([]int, 0)
	cnt := 0
	for i := range n {
		if arr[i] > t {
			cnt++
		} else {
			list = append(list, cnt)
			cnt = 0
		}
	}
	list = append(list, cnt)
	fmt.Print(slices.Max(list))
}