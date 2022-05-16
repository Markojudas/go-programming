package main

import (
	"fmt"
	"io"
	"os"
)

//TYPE Decoder
/*
	func(*Decoder) Decode
	func (dec *Decoder) Decode(v interface{}) error
		* godoc.org def:
			Decode reads the next JSON-encoded value from its input and stores it in the value pointed to by v.
			**Unmarshall (json.Unmarshall)

	Writer interface:
		* Package IO
		type Writer interface {
			Write(p []byte)(n int, err error)
		}
		***any TYPE that implements the method Write(p []byte)(n int, err error) IT IS ALSO implementing the Writer interface

		TYPE File implements the Writer interface since TYPE File has the method write

		**READ THE DOCS (fmt, io.Writer, os.Stdout)
*/

func main() {
	fmt.Fprintln(os.Stdout, "Hello World!") //this IS fmt.Println()

	io.WriteString(os.Stdout, "Hello from WriterString\n") //this IS also as fmt.Println()
}
