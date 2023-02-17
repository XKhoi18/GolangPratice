package main

import "fmt"

func main() {
	ch := make(chan int)
	go exChannel(ch)
	receive := <-ch
	fmt.Println(receive)
}

func exChannel(ch chan int) {
	sendValue := 5
	ch <- sendValue
}
