package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"math"
	"slices"
	"cmp"
)

type Point struct {
	id int
	x, y int
	d int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	points := make([]Point, n)
	for i := 0; i < n; i++ {
		points[i].id = i+1
		scanner.Scan()
		points[i].x, _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		points[i].y, _ = strconv.Atoi(scanner.Text())

		points[i].d = int(math.Abs(float64(points[i].x)) + math.Abs(float64(points[i].y)))
	}
	// Please write your code here.
	slices.SortFunc(points, func(a, b Point) int {
		// 1. 거리 기준으로 비교
		if n := cmp.Compare(a.d, b.d); n != 0 {
			return n
		}

		// 2. 거리가 같다면 id 기준으로 비교
		return cmp.Compare(a.id, b.id)
	})

	for _, v := range points {
		fmt.Println(v.id)
	}
}