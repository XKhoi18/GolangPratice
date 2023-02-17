package main

import (
	"fmt"
)

func main() {
	S1 := "ervervige"
	fmt.Println(solution(S1))
	S2 := "aaabab"
	fmt.Println(solution(S2))
	S3 := "x"
	fmt.Println(solution(S3))
}

func solution(S string) int {
	n := len(S)
	var fre [26]int
	for i := 0; i < n; i++ {
		fre[S[i]-'a'] += 1
	}
	var count int
	for j := 0; j < 26; j++ {
		if fre[j]%2 == 1 {
			count += 1
		}
	}
	if count == 0 || count == 1 {
		return 0
	} else {
		return count - 1
	}
}
