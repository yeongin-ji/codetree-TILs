package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	buf := make([]byte, 100_100)
	scanner.Buffer(buf, 100_100)

	scanner.Scan()
	word1 := scanner.Text()
	scanner.Scan()
	word2 := scanner.Text()
	
	

	// Please write your code here.
	if len(word1) != len(word2) {
		fmt.Println("No")
		return
	}

	r1 := []rune(word1)
	r2 := []rune(word2)
	slices.Sort(r1)
	slices.Sort(r2)
	word1 = string(r1)
	word2 = string(r2)
	if word1 == word2 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}