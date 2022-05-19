package main

/*
	Fix race condition using the atomic package!
*/

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	//to use the atomic package the "counter" needs to be of type int64

	var counter int64
	var wg sync.WaitGroup
	limit := 50

	wg.Add(limit)

	for i := 0; i < limit; i++ {
		go func() {
			//v := counter
			//increment the count
			atomic.AddInt64(&counter, 1)
			runtime.Gosched() //give CPU to another goroutine
			//v++
			//counter = v
			//fmt.Println("Counter thus far:", counter)
			//Reading the resource:
			fmt.Println("Counter thus far:", atomic.LoadInt64(&counter))
			wg.Done()
		}()
		//fmt.Println("Goroutines", runtime.NumGoroutine())
	}

	wg.Wait()

	fmt.Println("Counter:", counter)
}
