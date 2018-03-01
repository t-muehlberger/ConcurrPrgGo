package main

import (
	"fmt"
	"time"
)

func main() {
	tasks := []int{1, 2, 3, 4, 5}

	fmt.Println("==========================")
	fmt.Println("Sequential:")
	fmt.Println("==========================")

	start := time.Now()

	processSeq(tasks)

	fmt.Printf("Execution took: %s\n", time.Since(start))

	fmt.Println("==========================")
	fmt.Println("Pipeline:")
	fmt.Println("==========================")

	start = time.Now()

	processPipe(tasks)

	fmt.Printf("Execution took: %s\n", time.Since(start))
}

func processSeq(tasks []int) {
	for _, t := range tasks { // sequentielles ausführen aller Tasks
		aOut := doA(t)
		bOut := doB(aOut)
		res := doC(bOut)
		fmt.Printf("Task: %d ... Finished!\n", res)
	}
}

func processPipe(tasks []int) {
	// Kanäle zur Kommunikation zwischen Stufen
	in, aOut, bOut, res := make(chan int), make(chan int), make(chan int), make(chan int)

	go func() { // senden des Input an Pipeline
		for _, t := range tasks {
			in <- t
		}
		close(in)
	}()
	go stageA(in, aOut) // starten der Stufen
	go stageB(aOut, bOut)
	go stageC(bOut, res)

	// empfangen des Output der Pipeline
	for t := range res {
		fmt.Printf("Task: %d ... Finished!\n", t)
	}
}
func stageA(in <-chan int, out chan<- int) {
	for t := range in {
		out <- doA(t) // sequentielles Ausführen von doA für jedes Element in der Pipeline
	}
	close(out)
}

func stageB(in <-chan int, out chan<- int) {
	for t := range in {
		out <- doB(t)
	}
	close(out)
}

func stageC(in <-chan int, out chan<- int) {
	for t := range in {
		out <- doC(t)
	}
	close(out)
}

func doA(t int) int {
	time.Sleep(300 * time.Millisecond)
	fmt.Printf("Task: %d Stage A ... Done!\n", t)
	return t
}

func doB(t int) int {
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("Task: %d Stage B ... Done!\n", t)
	return t
}

func doC(t int) int {
	time.Sleep(200 * time.Millisecond)
	fmt.Printf("Task: %d Stage C ... Done!\n", t)
	return t
}
