package main

//instead of mutex, using atomic - gotta use int64

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())

	var counter int64

	var wg sync.WaitGroup
	const gs = 100
	wg.Add(gs)

	fmt.Println()
	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)                         //write to counter
			runtime.Gosched()                                    //allow something run.
			fmt.Println("Counter\t", atomic.LoadInt64(&counter)) //read counter
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println()

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())
	fmt.Println("Counter:", counter)
}
