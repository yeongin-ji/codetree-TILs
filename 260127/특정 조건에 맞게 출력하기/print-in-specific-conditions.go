package main

import "fmt"

func main() {
	var arr [110]int
	i := 0
	for {
		if _, err := fmt.Scan(&arr[i]); err != nil {
			break
		}
		i++
	}
	s := arr[:i]
	for _, v := range s {
		if v==0 {
			break
		}
		if v%2==1 {
			fmt.Printf("%d ", v+3)
		} else {
			fmt.Printf("%d ", v/2)
		}
	}
}