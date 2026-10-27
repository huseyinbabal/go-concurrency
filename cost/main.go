package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	const n = 100_000
	var before, after runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&before)

	block := make(chan struct{})
	start := time.Now()
	for range n {
		go func() { <-block }()
	}
	took := time.Since(start)
	runtime.ReadMemStats(&after)

	fmt.Println("goroutines:", runtime.NumGoroutine())
	fmt.Println("started in:", took.Round(time.Millisecond))
	fmt.Println("stack memory:", (after.StackInuse-before.StackInuse)/1024/1024, "MiB")
	fmt.Println("per goroutine:", (after.StackInuse-before.StackInuse)/n, "bytes of stack")
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))
	close(block)
}
