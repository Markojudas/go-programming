package main

/*
	part 1)
	Create a race condition (use the -race flag when running main.go)

	part2)
	Fix the race conditions using mutex
*/

import (
	"fmt"
	"runtime"
	"sync"
)

//for part 2 creating a mutex (lock) to fix race conditions
var lock sync.Mutex

func main() {

	counter := 0
	var wg sync.WaitGroup
	limit := 50

	wg.Add(limit)

	for i := 0; i < limit; i++ {
		go func() {
			//locking the critical section (part 2)
			lock.Lock()
			v := counter
			runtime.Gosched() //give CPU to another goroutine
			v++
			counter = v
			fmt.Println("Counter thus far:", counter)
			//no longer using the critical section - release the lock! (part 2)
			lock.Unlock()
			wg.Done()
		}()
		//fmt.Println("Goroutines", runtime.NumGoroutine())
	}

	wg.Wait()

	fmt.Println("Counter:", counter)
}
