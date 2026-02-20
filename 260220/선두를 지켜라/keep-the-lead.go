package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Move(seg, h [][2]int, id int) int {
	currT := 0
	currPos := 0
	for _, s := range seg {
		v := s[0]
		for range s[1] {
			currT++
			currPos += v
			h[currT][id] = currPos
		}
	}
	return currT
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	
	nextInt := func() int {
		scanner.Scan()
		num, _ := strconv.Atoi(scanner.Text())
		return num
	}
	
	n := nextInt()
	m := nextInt()
	
	aSegments := make([][2]int, n)
	for i := 0; i < n; i++ {
		aSegments[i][0] = nextInt()
		aSegments[i][1] = nextInt()
	}
	
	bSegments := make([][2]int, m)
	for i := 0; i < m; i++ {
		bSegments[i][0] = nextInt()
		bSegments[i][1] = nextInt()
	}
	
	// Please write your code here.
	var totalT int
	hist := make([][2]int, 1_000_000*n+100)
	_ = Move(aSegments, hist, 0)
	totalT = Move(bSegments, hist, 1)

	head := make([]int, 0)
	for i := 2; i <= totalT; i++ {
		offset := hist[i][0]-hist[i][1]
		if offset < 0 { // b가 선두
			head = append(head, 1)
		} else if offset > 0 { // a가 선두
			head = append(head, 0)
		}
	}

	// 연속부분수열 문제로 바뀜
	cnt := 0
	for i := range len(head) {
		if i!=0 && head[i]!=head[i-1] {
			cnt++
		}
	}

	fmt.Print(cnt)
}