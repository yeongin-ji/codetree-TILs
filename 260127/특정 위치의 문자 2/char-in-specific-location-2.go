package main

import "fmt"

func main() {
	var arr [20]string
	for i := range 10 {
		fmt.Scan(&arr[i])
	}
	fmt.Printf("%s %s %s", arr[1], arr[4], arr[7])
}