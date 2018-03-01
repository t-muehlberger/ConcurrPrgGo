package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		usage()
	}
	num, err := strconv.Atoi(os.Args[1])
	if err != nil || num <= 0 {
		usage()
	}

	started := make(chan bool)
	canExit := make(chan bool)

	for i := 0; i < num; i++ {
		go func() {
			started <- true
			<-canExit // block the goroutine
		}()
	}

	for i := 0; i < num; i++ {
		<-started
	}

	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)

	fmt.Printf("%d goroutines are started.", num)
	fmt.Println("Press ENTER to exit")
	bufio.NewScanner(os.Stdin).Scan()
}

func usage() {
	fmt.Println("Starts a given number of goroutines and keeps them in a blocked state")
	fmt.Printf("Usage: %s <n>\n", os.Args[0])
	fmt.Println("n ... A positive number greater than 0")
	os.Exit(1)
}
