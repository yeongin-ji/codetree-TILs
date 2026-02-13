package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func sort(nums []int) {
	for i := 1; i < len(nums); i++ {
		for j := i; j > 0; j-- {
			if nums[j] < nums[j-1] {
				swap(nums, j, j-1)
			} else {
				break
			}
		}
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