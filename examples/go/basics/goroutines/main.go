// This program demonstrates how to quickly start
// a goroutine and also how to profile memory usage.
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	w := &sync.WaitGroup{}

	// Go v1.25+
	for range 1000000 {
		w.Go(func() {
			time.Sleep(1 * time.Second)
		})
	}

	w.Wait()
}

// For versions < v1.25
func _() {
	mem := &runtime.MemStats{}
	runtime.ReadMemStats(mem)
	fmt.Printf("Initial memory usage: %v KB\n", mem.Alloc/1024)

	w := &sync.WaitGroup{}
	for range 1000000 {
		w.Add(1)
		go (func() {
			defer w.Done()
			time.Sleep(1 * time.Second)
		})()
	}

	runtime.ReadMemStats(mem)
	fmt.Printf("Memory usage after allocation: %v KB\n", mem.Alloc/1024)
	runtime.GC()
	runtime.ReadMemStats(mem)
	fmt.Printf("Memory usage after garbage collection: %v KB\n", mem.Alloc/1024)
}
