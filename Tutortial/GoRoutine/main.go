package main

import (
	"fmt"
	"time"
)

func main() {
	go go1()
	go go2()
	time.Sleep(time.Second) // sleep 1 second to wait for go1 and go2 finish
}

func go1() {
	for i := 1; i < 10; i++ {
		go fmt.Println(i)
	}
}

func go2() {
	for i := 1; i < 20; i++ {
		go fmt.Println(i)
	}
}
