package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func f(a, b *int) {
	var l, s *int
	if *a < *b {
		l, s = b, a
	} else {
		l, s = a, b
	}
	*s = *s + 10
	*l = *l * 2
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	a, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	b, _ := strconv.Atoi(scanner.Text())

	// Please write your code here.
	f(&a, &b)
	fmt.Print(a, b)
}