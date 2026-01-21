package main

import "fmt"

func main() {
	m := map[int]string{
		1: "John",
		2: "Tom",
		3: "Paul",
		4: "Sam",
	}	
	for {
		var n int
		fmt.Scanf("%d", &n)
		if v := m[n]; v == "" {
			fmt.Println("Vacancy")
			break
		} else {
			fmt.Println(v)
		}
	}
}