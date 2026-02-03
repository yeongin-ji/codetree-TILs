package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	if strings.Contains(s, "ee") {
		fmt.Print("Yes ")
	} else {
		fmt.Print("No ")
	}
	if strings.Contains(s, "ab") {
		fmt.Print("Yes ")
	} else {
		fmt.Print("No ")
	}
}