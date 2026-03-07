package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var grid, visited [][]int
var n, m int
var ds [2][2]int

func DFS(x, y int) {
	visited[x][y] = 1

	if x==n && y==m {
		return
	}

	for _, d := range ds {
		dx, dy := d[0], d[1]
		newx := x + dx
		newy := y + dy

		if IsSafe(newx, newy) {
			DFS(newx, newy)
		}
	}
}

func IsSafe(x, y int) bool {
	if !InRange(x, y) {
		return false
	}
	if grid[x][y] == 0 {
		return false
	}
	if visited[x][y] == 1 {
		return false
	}
	return true
}

func InRange(x, y int) bool {
	if x <= n && y <= m {
		return true
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ = strconv.Atoi(scanner.Text())

	grid = make([][]int, n+1)
	visited = make([][]int, n+1)
	for i := 1; i <= n; i++ {
		grid[i] = make([]int, m+1)
		visited[i] = make([]int, m+1)
		for j := 1; j <= m; j++ {
			scanner.Scan()
			grid[i][j], _ = strconv.Atoi(scanner.Text())
		}
	}

	// Please write your code here.
	ds = [2][2]int{{1, 0}, {0, 1}}
	DFS(1, 1)

	if visited[n][m] != 0 {
		fmt.Print(1)
	} else {
		fmt.Print(0)
	}
}