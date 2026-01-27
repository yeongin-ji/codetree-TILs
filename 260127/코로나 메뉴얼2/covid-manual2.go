package main

import "fmt"

func main() {
	m := [5]struct {
		yn string
		fv int
	}{}
	cnt := map[string]int{
		"A": 0,
		"B": 0,
		"C": 0,
		"D": 0,
	}
	for i := range 3 {
		fmt.Scan(&m[i].yn, &m[i].fv)
	}
	for i := range 3 {
		if m[i].yn=="Y" {
			if m[i].fv >= 37 {
				cnt["A"]++
			} else {
				cnt["C"]++
			}
		} else if m[i].yn=="N" {
			if m[i].fv >= 37 {
				cnt["B"]++
			} else {
				cnt["D"]++
			}
		}
	}
	var e string
	if cnt["A"] >= 2 {
		e = "E"
	} 
	fmt.Print(cnt["A"], cnt["B"], cnt["C"], cnt["D"])
	fmt.Printf(" %s", e)
}