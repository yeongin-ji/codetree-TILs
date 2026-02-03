package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	var n int
	fmt.Scan(&n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	var con string
	for range n {
		scanner.Scan()
		con += scanner.Text()
	}	
	for i, r := range con {
		if i!=0 && i%5==0 {
			fmt.Println()
		}
		fmt.Printf("%c", r)
	}
}