package main

import "fmt"

func main() {
	var arr [2][4]int
	for i := range 2 {
		for j := range 4 {
			fmt.Scan(&arr[i][j])
		}
	}	
	var rst float64
	for i := range 2 {
		sum := 0
		for j := range 4 {
			sum += arr[i][j]
		}
		avg := float64(sum)/4
		rst += avg
		fmt.Printf("%.1f ", avg)
	}
	fmt.Println()
	for i := range 4 {
		sum := 0
		for j := range 2 {
			sum += arr[j][i]
		}
		avg := float64(sum)/2
		fmt.Printf("%.1f ", avg)
	}
	fmt.Println()
	rst /= 2
	fmt.Printf("%.1f", rst)
}