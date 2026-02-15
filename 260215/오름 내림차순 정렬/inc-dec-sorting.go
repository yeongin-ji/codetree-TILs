package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func merge_sort(nums []int) {
	//fmt.Printf("merge_sort(%d)\n", len(nums))
	// 종료 조건
	if len(nums) == 1 {
		//fmt.Println("len(nums)==1, return")
		return
	}

	// divide and conpuer
	n1, n2 := nums[:len(nums)/2], nums[len(nums)/2:]
	merge_sort(n1)
	merge_sort(n2)

	// combine
	tmp := make([]int, len(nums))
	var c1, c2 int
	for i := range len(nums) {
		//fmt.Printf("loop %d, c1:%d, c2:%d\n", i+1, c1, c2)
		// 한쪽 인덱스가 이미 끝에 다다랐다면 다른쪽 원소를 가져오기
		if c1 == len(nums)/2 {
			tmp[i] = n2[c2]
			c2++
			continue
		}
		if c2 == len(nums)/2+len(nums)%2 {
			tmp[i] = n1[c1]
			c1++
			continue
		}

		// 양쪽 원소를 비교하고 인덱스 조정
		if n1[c1] <= n2[c2] {
			tmp[i] = n1[c1]
			c1++
		} else {
			tmp[i] = n2[c2]
			c2++
		}
		//fmt.Printf("loop %d, c1:%d, c2:%d\n\n", i+1, c1, c2)
	}
	copy(nums, tmp)
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
	merge_sort(nums)
	for _, v := range nums {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
	r_nums := reverse(nums)
	for _, v := range r_nums {
		fmt.Printf("%d ", v)
	}
}