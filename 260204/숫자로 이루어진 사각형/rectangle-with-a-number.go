package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func f(n int) {
	cnt := 1
	for range n {
		for range n {
			fmt.Printf("%d ", cnt)
			cnt = cnt%9 + 1
		}
		fmt.Println()
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	
	// Please write your code here.
	f(n)
}