package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	s := scanner.Text()
	scanner.Scan()
	c := scanner.Text()
	flag := false
	for i, r := range s {
		if r == rune(c[0]) {
			fmt.Print(i)
			flag = true
			break
		}
	}
	if !flag {
		fmt.Print("No")
	}
}