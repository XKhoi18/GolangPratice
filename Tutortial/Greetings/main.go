package main

import (
	"example/greetings/greetings"
	"fmt"
)

func main() {
	name := "David"
	fmt.Println(greetings.Hello(name))
}
