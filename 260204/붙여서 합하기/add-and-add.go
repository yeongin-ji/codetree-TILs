package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b string
	fmt.Scanf("%s %s", &a, &b)
	ab, ba := a+b, b+a
	ai, _ := strconv.Atoi(ab)
	bi, _ := strconv.Atoi(ba)
	fmt.Print(ai+bi)
}