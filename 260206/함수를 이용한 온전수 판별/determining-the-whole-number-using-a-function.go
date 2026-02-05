package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func isComplete(n int) bool {
	if n%2==0 {
		return false
	} else if n%10==5 {
		return false
	} else if n%3==0 && n%9!=0 {
		return false
	} else {
		return true
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	a, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	b, _ := strconv.Atoi(scanner.Text())

	// Please write your code here.
	cnt := 0
	for i := a; i <= b; i++ {
		if isComplete(i) {
			cnt++
		}
	}
	fmt.Print(cnt)
}