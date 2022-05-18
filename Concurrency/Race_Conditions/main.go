package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())

	counter := 0

	var wg sync.WaitGroup
	const gs = 100
	wg.Add(gs)

	var mu sync.Mutex

	fmt.Println()
	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock() //locking the critical section so noone can access it
			v := counter
			runtime.Gosched() //allow something run.
			v++
			counter = v
			mu.Unlock() //unlocking the critical section
			wg.Done()
		}()
		fmt.Println("Goroutines", runtime.NumGoroutine())

	}

	wg.Wait()
	fmt.Println()

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())
	fmt.Println("Counter:", counter)
}
