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
	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	t := scanner.Text()

	words := make([]string, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		words[i] = scanner.Text()
	}

	// Please write your code here.
	tmp := make([]string, 0)
	for _, s := range words {
		r := []rune(s)
		if len(t) <= len(r) && t == string(r[:len(t)]) {
			tmp = append(tmp, s)
		}
	}
	slices.Sort(tmp)
	for i, s := range tmp {
		if i+1==k {
			fmt.Println(s)
		}
	}
}