package main

/*
	Encode to JSON and print to Stdout

	*Keep in mind:
		*	func (enc *Encoder) Encode(v any) error  -> signature; we need TYPE Encoder pointer:
		*   to get a TYPE Encodder pointer:  func NewEncoder(w io.Writer) *Encoder
		*   We need to pass a value of type io.Writer
			* io.Writer is an interface, so we need a type that implements it = TYPE File (os.File) who has the Write(b []byte) (n int, err error) method
			* func (f *File) Write(b []byte) (n int, err error)
			* os.Stdout is a variable for the package OS = Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout"),
				* NewFile is a function: func NewFile(fd uintptr, name string) *File; that returns a TYPE *File (pointer to a file)\
	* json.NewEncoder(os.Stdout) -> *Encoder.Encode(any variable)
	* json.NewEncoder(os.Stdout).Encode(users) = this will encode and write the output to stdout (to screen)
		* since the Encode function returns an error we have to assign it to a variable and you always want to deal with the variable to holds the "return error"
		* err := json.NewEncoder(os.Stdout).Encode(users)
*/

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	u1 := User{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := User{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := User{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []User{u1, u2, u3}

	fmt.Println(users)
	fmt.Println()

	err := json.NewEncoder(os.Stdout).Encode(users)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println()
	}

}
