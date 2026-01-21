package main

import "fmt"

func main() {
	var n float64
	fmt.Scanf("%f", &n)
	i := 1
	cnt := 0
	for {
		n /= float64(i)
		cnt++
		if n <= 1 {
			break
		}
		i++
	}	
	fmt.Print(cnt)
}