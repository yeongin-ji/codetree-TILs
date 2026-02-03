package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	var arr [10]string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := range 10 {
		scanner.Scan()
		arr[i] = scanner.Text()
	}	
	for i := range 10 {
		fmt.Println(arr[9-i])
	}
}