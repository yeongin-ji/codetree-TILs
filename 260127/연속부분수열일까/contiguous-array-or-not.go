package main

import "fmt"

func main() {
	var N1, N2 int
	var a, b [110]int
	flag := false
	fmt.Scan(&N1, &N2)
	for i := range N1 {
		fmt.Scan(&a[i])
	}	
	for i := range N2 {
		fmt.Scan(&b[i])
	}
	for i := range N1 {
		found := true
		for j := i; j < i+N2; j++ {
			if j >= N1-1 || a[j]!=b[j-i] {
				found = false
				break
			}
		}
		if found {
			flag = true
			break
		}
	}
	if flag {
		fmt.Print("Yes")
	} else {
		fmt.Print("No")
	}
}