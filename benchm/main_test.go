package main

import (
	"testing"
)

func BenchmarkGoroutine(b *testing.B) {
	started := make(chan bool)
	for n := 0; n < b.N; n++ {
		go func() { started <- true }()
		<-started // Warten, bis die goroutine gestartet wurde
	}
}
