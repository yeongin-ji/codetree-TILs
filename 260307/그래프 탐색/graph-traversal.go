package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var adj [][]int
var visited []int

func dfs(v, n int) {
	for i := 1; i <= n; i++ {
		curr := i
		if adj[v][curr]==1 && visited[curr]==0 {
			visited[curr] = 1
			dfs(curr, n)
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	edgesX := make([]int, m+1)
	edgesY := make([]int, m+1)
	for i := 1; i <= m; i++ {
		scanner.Scan()
		x, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		y, _ := strconv.Atoi(scanner.Text())
		edgesX[i] = x
		edgesY[i] = y
	}

	// Please write your code here.
	visited = make([]int, n+1)
	adj = make([][]int, n+1)
	for i := 1; i <= n; i++ {
		adj[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		adj[edgesX[i]][edgesY[i]] = 1
		adj[edgesY[i]][edgesX[i]] = 1
	}

	dfs(1, n)

	var cnt int
	for i := 1; i <= n; i++ {
		if visited[i] == 1 {
			cnt++
		}
	}
	fmt.Print(cnt-1)
}