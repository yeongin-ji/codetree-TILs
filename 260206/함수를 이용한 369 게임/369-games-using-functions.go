package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func isMagicNum(n int) bool {
	return is369(n) || n%3==0
}

func is369(n int) bool {
	s := strconv.Itoa(n)
	flag := false
	for _, c := range s {
		if c=='3' || c=='6' || c=='9' {
			flag = true
		}
	}
	return flag
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
		if isMagicNum(i) {
			cnt++
		}
	}
	fmt.Print(cnt)
}