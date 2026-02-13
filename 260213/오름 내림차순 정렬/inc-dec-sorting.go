package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// selection sort
func sort(nums []int) {
	for i := range len(nums)-1 {
		min_idx := i
		for j := i; j < len(nums); j++ {
			if nums[min_idx] > nums[j] {
				min_idx = j
			}
		}
		swap(nums, i, min_idx)
	}
}

func swap(nums []int, x, y int) {
	tmp := nums[x]
	nums[x] = nums[y]
	nums[y] = tmp
}

func reverse(nums []int) []int {
	rst := make([]int, len(nums))
	for i, v := range nums {
		rst[len(rst)-1-i] = v
	}
	return rst
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		nums[i], _ = strconv.Atoi(scanner.Text())
	}

	// Please write your code here.
	sort(nums)
	for _, v := range nums {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
	r_nums := reverse(nums)
	for _, v := range r_nums {
		fmt.Printf("%d ", v)
	}
}