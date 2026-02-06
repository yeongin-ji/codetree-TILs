package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func isPartial(a, b []int) bool {
	for i := range len(a) {
		found := true
		for j := range len(b) {
			if a[i+j] != b[j] {
				found = false
				break
			}
		}
		if found {
			return true
		}
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n1, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	n2, _ := strconv.Atoi(scanner.Text())

	a := make([]int, n1)
	for i := 0; i < n1; i++ {
		scanner.Scan()
		a[i], _ = strconv.Atoi(scanner.Text())
	}

	b := make([]int, n2)
	for i := 0; i < n2; i++ {
		scanner.Scan()
		b[i], _ = strconv.Atoi(scanner.Text())
	}

	// Please write your code here.
	if isPartial(a, b) {
		fmt.Print("Yes")
	} else {
		fmt.Print("No")
	}
}