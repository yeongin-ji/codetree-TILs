package main

import "fmt"

func main() {
	var cnt [10]int	
	var arr [110]int
	i := 0
	for {
		if _, err := fmt.Scan(&arr[i]); err != nil {
			break
		}
		i++
	}
	for j := range i {
		if arr[j]==0 {
			break
		}
		cnt[arr[j]/10]++
	}
	for i := 1; i <= 9; i++ {
		fmt.Printf("%d - %d\n", i, cnt[i])
	}
}