package main

import "fmt"

func main() {
	var arr [10]int
	for i := range 10 {
		fmt.Scan(&arr[i])
	}	
	evensum := 0
	var s, cnt int
	for i, v := range arr {
		if i%2!=0 {
			evensum += v
		}
		if (i+1)%3==0 {
			s += v
			cnt++
		}
	}
	avg := float64(s)/float64(cnt)
	fmt.Printf("%d %.1f", evensum, avg)
}