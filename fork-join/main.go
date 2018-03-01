package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Starting")
	answerChan := calcAnswer()
	fmt.Println("Doing something else")
	answer := <-answerChan
	fmt.Printf("The answer is: %d\n", answer)
}
func calcAnswer() <-chan int {
	res := make(chan int)
	go func() {
		time.Sleep(1000) // finden der Antwort benötigt viel Zeit
		res <- 42
	}()
	return res
}
