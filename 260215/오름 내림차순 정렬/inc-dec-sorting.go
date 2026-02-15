package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func quick_sort(nums []int) {
	// base condition, length가 0일 수도 있음
	if len(nums) <= 1 {
		return
	}

	// pivot 옮기기
	var p int
	s, e := 1, len(nums)-1
	for {
		for s <= e && nums[s] <= nums[p] {
			s++
		}
		for s <= e && nums[e] >= nums[p] {
			e--
		}
		if s > e {
			break
		} else {
			swap(nums, s, e)
		}
	}
	swap(nums, p, e)

	// pivot 자리 제외하고 양쪽 각각 정렬
	quick_sort(nums[:e])
	quick_sort(nums[e+1:])
}

func reverse(nums []int) []int {
	rst := make([]int, len(nums))
	for i, v := range nums {
		rst[len(rst)-1-i] = v
	}
	return rst
}

func swap(nums []int, x, y int) {
	tmp := nums[x]
	nums[x] = nums[y]
	nums[y] = tmp
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
	quick_sort(nums)
	for _, v := range nums {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
	r_nums := reverse(nums)
	for _, v := range r_nums {
		fmt.Printf("%d ", v)
	}
}