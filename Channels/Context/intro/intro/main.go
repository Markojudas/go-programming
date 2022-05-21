package main

import (
	"context"
	"fmt"
)

/*
	Read the docs!!!
	https://pkg.go.dev/context

*/

func main() {

	/*
		func Background() Context
		returns a non-nil empty Context. it is never canceled, has no values, and has no deadline
		it is typically used by the main function, initialization, and tests, and as the top-lvel Context for incoming requests
	*/
	ctx := context.Background()

	fmt.Println("Context:\t", ctx)

	/*
		Err() error
		If Done is not yet closed, Err returns nil.
		If Done is closed, Err returns a non-nil error explaining why:
		Cancelled if the context was cancelled
		Or deadlineExceeded if the context's deadline passed
		After Err returns a non-nil error, sucessive calls to err return the same error
	*/
	fmt.Println("context err:\t", ctx.Err())
	fmt.Printf("Context Type:\t%T\n", ctx) // *context.emptyCtx
	fmt.Println("******************************************************")

}
