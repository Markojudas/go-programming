package main

//concurrency

import (
	"fmt"
	"runtime"
	"sync"
)

//create a waitgroup variable from the sync()
var wg sync.WaitGroup

/*
	we are using sync.WaitGroup.Add(),Wait(),Done()
	the documentation shows that this methods have a receiver of type *WaitGroup, a pointer.

	the variable wg is TPE sync.WaitGroup and not pointer sync.WaitGroup but it is still work!!

	***method sets: The method set of a type determines the INTERFCES that type implements.

	therefore we are just calling the method and not using an interface.
*/

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)                 //shows OS
	fmt.Println("ARCH\t\t", runtime.GOARCH)             //shows architecture
	fmt.Println("CPUs\t\t", runtime.NumCPU())           //shows # of CPus (this is a function)
	fmt.Println("Goroutines\t", runtime.NumGoroutine()) //shows goroutines (this is a function) (this is 1)

	wg.Add((1))
	go foo() //own goroutine -> think of a fork on the road; different timeline!
	bar()

	fmt.Println("CPUs\t\t", runtime.NumCPU())           //same num of CPUS
	fmt.Println("Goroutines\t", runtime.NumGoroutine()) //since we sent foo() to its own go routine, now there are 2
	wg.Wait()                                           //wait untl the 1 (or more) goroutine are done

	/*
		without the WG variable the main() just finishes killing everything even goroutines that haven't "executed".
		the wg we add the number of what we have to wait for
		then when we mark where we want to wait() for that goroutine to execute
		once the goroutine is done (done()) the main() continues
	*/
}

func foo() {

	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done() //advise main() that this is done.
}

func bar() {

	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
