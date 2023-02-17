package main

import (
	"fmt"
	"sort"
)

func main() {
	A1 := []int{4,1,3}
	fmt.Println(Solution(A1))
}

func Solution(A []int) int {
	X := len(A)
	sort.Ints(A)
	for i:=0; i<X; i++{
		if A[i] != (i+1){
			return 0
		}
	}
	return 1
}

//cart = append(cart[0:index], cart[index+1:]...)
