package main

import "fmt"

func main() {
	for i := range 19 {
		for j := range 19 {
			fmt.Printf("%d * %d = %d ", i+1, j+1, (i+1)*(j+1))
			if j+1==19 || j%2!=0 {
				fmt.Println()
			} else {
				fmt.Print("/ ")
			}
		}
	}	
}