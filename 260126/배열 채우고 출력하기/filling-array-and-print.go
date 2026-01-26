package main

import "fmt"

func main() {
	var arr [10]string
	for i := range 10 {
		fmt.Scan(&arr[i])
	}	
	for i := 9; i >= 0; i-- {
        fmt.Print(arr[i])
    }
}