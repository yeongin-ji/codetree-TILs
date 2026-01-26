package main

import "fmt"

func main() {
	var n, cnt int
	fmt.Scan(&n)
	for range n {
		var score [4]int
		sum := 0
		for i := range 4 {
			fmt.Scan(&score[i])
		}
		for _, v := range score {
			sum += v
		}
		avg := float64(sum)/4
		if avg >= 60 {
			fmt.Println("pass")
			cnt++
		} else {
			fmt.Println("fail")
		}
	}	
	fmt.Print(cnt)
}