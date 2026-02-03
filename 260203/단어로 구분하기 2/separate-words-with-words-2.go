package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	cnt := 1
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		if cnt%2==1 {
			fmt.Println(scanner.Text())
		}
		cnt++
	}
}