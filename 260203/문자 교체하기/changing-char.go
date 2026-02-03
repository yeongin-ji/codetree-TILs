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
	s1 := scanner.Text()
	scanner.Scan()
	s2 := scanner.Text()
	r1 := []rune(s1)
	r2 := []rune(s2)
	r2[0] = r1[0]
	r2[1] = r1[1]
	s2 = string(r2)
	fmt.Print(s2)
}