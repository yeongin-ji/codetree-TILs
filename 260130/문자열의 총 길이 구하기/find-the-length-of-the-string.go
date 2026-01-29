package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	cnt := 0
	for scanner.Scan() {
		s := scanner.Text()
		l := len(s)
		cnt += l
	}
	fmt.Print(cnt)
}