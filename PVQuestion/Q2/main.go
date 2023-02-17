package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	message1 := "And now here is my secret"
	fmt.Println(Solution(message1))
	message2 := "There is an animal with four legs"
	fmt.Println(Solution(message2))
}

func Solution(message string) string {
	maxLenght := 15
	if len(message) <= 15 {
		return message
	} else {
		trimpMessage := string(message[0 : maxLenght-3])
		trimpMessage = string(trimpMessage[0:int64(math.Min(float64(len(trimpMessage)), float64(strings.LastIndex(trimpMessage, " "))))])
		return trimpMessage + "..."
	}
}
