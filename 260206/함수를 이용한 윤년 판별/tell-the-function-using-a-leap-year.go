package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func isYoon(y int) bool {
	if y%4!=0 {
		return false
	}
	if y%100==0 && y%400!=0 {
		return false
	}
	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	y, _ := strconv.Atoi(input)

	// Please write your code here.
	if isYoon(y) {
		fmt.Print("true")
	} else {
		fmt.Print("false")
	}
}