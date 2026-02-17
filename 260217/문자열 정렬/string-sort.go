package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()
	// Please write your code here.

	r := []rune(str)
	slices.Sort(r)
	str = string(r)
	fmt.Print(str)
}