package main

import (
	"fmt"
	"runtime"
	"time"
)

func getMemoryUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Memory Usage: %v MiB\n", m.Alloc/1024/1024)
}

func main() {
	// Print memory usage before any allocations
	getMemoryUsage()

	// Allocate some memory
	for i := 0; i < 10; i++ {
		s := make([]byte, 1<<20) // Allocate 1 MiB
		_ = s
		time.Sleep(time.Second)
		getMemoryUsage()
	}
}
