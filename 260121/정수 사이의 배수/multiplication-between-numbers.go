package main

import "fmt"

func main() {
	var a, b, cnt, sum int
	var avg float64
	fmt.Scanf("%d %d", &a, &b)
	for i := a; i <= b; i++ {
		if i%5==0 || i%7==0 {
			sum += i
			cnt++
		}
	}
	avg = float64(sum)/float64(cnt)
	fmt.Printf("%d %.1f", sum, avg)
}