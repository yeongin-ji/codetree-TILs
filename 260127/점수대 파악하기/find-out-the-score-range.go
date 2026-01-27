package main

import "fmt"

func main() {
	var arr [110]int
	var cnt [11]int
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
	for i := 10; i >= 1; i-- {
		fmt.Printf("%d - %d\n", i*10, cnt[i])
	}
}