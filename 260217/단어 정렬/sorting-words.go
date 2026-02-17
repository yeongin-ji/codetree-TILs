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

	words := make([]string, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		words[i] = scanner.Text()
	}

	// Please write your code here.
	slices.Sort(words)
	for _, v := range words {
		fmt.Println(v)
	}
}