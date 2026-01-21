package main

import "fmt"

func main() {
	var n [5]int
	cnt := 0
	for i := range 5 {
		fmt.Scanf("%d", &n[i])
		if n[i]%2==0 {
			cnt++
		}
	}	
	fmt.Print(cnt)

}