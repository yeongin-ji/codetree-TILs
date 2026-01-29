package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 표준 입력(os.Stdin)을 읽는 스캐너 생성
	scanner := bufio.NewScanner(os.Stdin)
	
	// 한 줄을 읽음 (성공하면 true 반환)
	var line string
	if scanner.Scan() {
		line = scanner.Text() // 읽은 내용을 문자열로 가져옴
	}
	i := 1
	for _, c := range line {
		if i >= 3 && i <= 10 {
			fmt.Printf("%c", c)
		}
		i++
	}
}