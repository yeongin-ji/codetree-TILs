package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()

	// Please write your code here.
	r := []rune(str)
	cnt := 0
	for i := range len(r) {
		if r[i]=='(' {
			for j := i+1; j < len(r); j++ {
				if r[j] == ')' {
					cnt++
				}
			}
		}
	}
	fmt.Print(cnt)
}