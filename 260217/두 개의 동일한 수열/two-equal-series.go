package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"slices"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	a := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		a[i], _ = strconv.Atoi(scanner.Text())
	}

	b := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		b[i], _ = strconv.Atoi(scanner.Text())
	}

	// Please write your code here.
	slices.Sort(a)
	slices.Sort(b)
	if slices.Equal(a, b) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}