package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	cnt := 0
	i := 1
	for {
		fmt.Printf("%d ", n*i)
		if (n*i)%5==0 {
			cnt++
		}
		if cnt == 2 {
			break
		}
		i++
	}	
}