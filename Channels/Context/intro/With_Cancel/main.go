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

	fmt.Println("\nWITH CANCEL")

	/*
		func WithCancel
		func WithCancel(parent Context) (ctx Context, cancel CancelFunc) -> takes in a Context and returns another context AND the cancel function

		WithCancel returns a copy of parent with a new Done channel. The return context's Done channel is closed when the return cancel function is called
		or when the parent context's Done channel is closed, whichever happens first.

		Cancelling this context releases resources associated with it, so code should call cancel as soon as the operations running in this Context complete
	*/

	fmt.Println()
	ctx, cnxFunc := context.WithCancel(ctx)
	fmt.Println("Context:\t", ctx)           //context.Background.WithCancel
	fmt.Println("context err:\t", ctx.Err()) //nil
	fmt.Printf("Context Type:\t%T\n", ctx)   // *context.cancelCtx
	fmt.Println("******************************************************")

	fmt.Println("\nCALLING CANCEL")
	fmt.Println()
	cnxFunc()

	fmt.Println("Context:\t", ctx)
	fmt.Println("Context err:\t", ctx.Err())  //context canceled
	fmt.Printf("Context Type:\t%T\n", ctx)    //*conext.cancelCtx
	fmt.Println("Cancel:\t\t", cnxFunc)       //address for the cancelFunc
	fmt.Printf("Cancel Type:\t%T\n", cnxFunc) //context.CancelFunc

}
