package main

import "fmt"

func main() {
	var N, Q int
	var arr [110]int
	fmt.Scan(&N, &Q)
	for i := range N {
		fmt.Scan(&arr[i])
	}
	for range Q {
		var cmd int
		fmt.Scan(&cmd)
		switch cmd {
		case 1:
			var a int
			fmt.Scan(&a)
			fmt.Println(arr[a-1])
		case 2:
			var b, idx int
			fmt.Scan(&b)
			for j := range N {
				if arr[j]==b {
					idx = j+1
					break
				}
			}
			fmt.Println(idx)
		case 3:
			var s, e int
			fmt.Scan(&s, &e)
			for j := s; j <= e; j++ {
				fmt.Printf("%d ", arr[j-1])
			}
			fmt.Println()
		}
	}
}