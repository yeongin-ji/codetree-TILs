package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func f(m, d int) bool {
	if m > 12 || d > 31  {
		return false
	} 
	switch m {
	case 4, 6, 9, 11:
		if d > 30 {
			return false
		} else {
			return true
		}
	case 2:
		if d > 28 {
			return false
		} else {
			return true
		}
	}
	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	d, _ := strconv.Atoi(scanner.Text())

	// Please write your code here.
	if f(m, d) {
		fmt.Print("Yes")
	} else {
		fmt.Print("No")
	}
}