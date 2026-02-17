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

	nums := make([]int, 2*n)
	for i := 0; i < 2*n; i++ {
		scanner.Scan()
		num, _ := strconv.Atoi(scanner.Text())
		nums[i] = num
	}

	// Please write your code here.
	slices.Sort(nums)
	s := make([]int, n)
	for i := range n {
		s[i] = nums[i] + nums[len(nums)-1-i]
	}
	rst := slices.Max(s)
	fmt.Println(rst)
}