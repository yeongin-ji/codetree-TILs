package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	binary := scanner.Text()
	// Please write your code here.

	var rst int
	for _, b := range binary {
		rst = rst*2 + int(b-'0')
	}
	fmt.Print(rst)
}