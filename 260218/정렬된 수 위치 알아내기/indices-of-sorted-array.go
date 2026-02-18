package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"slices"
	"cmp"
)

type Index struct {
	idx int
	val int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	// 입력 단계
	var index = make([]Index, n)
	for i := 0; i < n; i++ {
		index[i].idx = i+1
		scanner.Scan()
		index[i].val, _ = strconv.Atoi(scanner.Text())
	}

	// 정렬 수행
	slices.SortFunc(index, func(a, b Index) int {
		if n := cmp.Compare(a.val, b.val); n != 0 {
			return n
		}
		return cmp.Compare(a.idx, b.idx)
	})

	// 본래 id 서치
	for i := range len(index) {
		for j, v := range index {
			if i+1 == v.idx {
				fmt.Printf("%d ", j+1)
			}
		}
	}
}