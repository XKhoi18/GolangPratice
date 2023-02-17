package main

import "fmt"

func main() {
	input := "Today is a new day, a better day."
	reverse := Reverse(input)
	doubleReverse := Reverse(reverse)
	fmt.Printf("original: %v \n", input)
	fmt.Printf("Reverse: %v \n", reverse)
	fmt.Printf("Reverse again: %v \n", doubleReverse)
}

func Reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}
