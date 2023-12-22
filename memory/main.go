package main

import (
	"fmt"
	"runtime"
)

func main() {
	numcpu := runtime.NumCPU()
	fmt.Printf("Number of cpu %d\n", numcpu)
}
