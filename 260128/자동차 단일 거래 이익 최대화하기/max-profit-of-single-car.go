package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var arr [1100]int
	for i := range n {
		fmt.Scan(&arr[i])
	}	
	m := map[int]int{}
	k := 0
	for i := range n {
		for j := i+1; j < n; j++ {
			m[k] = arr[j]-arr[i]
			k++
		}
	}
	prof := 0
	for _, v := range m {
		if v > prof {
			prof = v
		}
	}
	fmt.Print(prof)
}