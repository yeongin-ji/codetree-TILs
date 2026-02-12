package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func f(arr []int) int {
	if len(arr)==1 {
		return arr[0]
	}
	if len(arr)==2 {
		return lcm(arr[0], arr[1])
	} else {
		return lcm(f(arr[:len(arr)-1]), arr[len(arr)-1])
	}
}

func lcm(a, b int) int {
	return (a*b) / gcd(a, b)
}

func gcd(a, b int) int {
	gcd := 1
	for i := 1; i <= a; i++ {
		if a%i==0 && b%i==0 {
			gcd = i
		}
	}
	return gcd
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		arr[i], _ = strconv.Atoi(scanner.Text())
	}

	// Please write your code here.
	rst := f(arr)
	fmt.Println(rst)
}