/*
Defer is used to ensure that a function call is perform later in a program's execution,
usually for purposes of cleanup. `defer` is often used where e.g. `ensure` and `finally`
would be used in other languages
*/

package main

import (
	"fmt"
	"os"
)

func main() {

	/*
		Suppose we wanted to crrate a file, write to it, and then close when we are done.
		Here is how we could that with `defer

		Immediately after getting a file object with `createFile`, we defer the closing of that file with `closeFile`.
		This will be executed at the end of the enclosing function (main), after `writeFile` has finished.
	*/

	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

// It is important to check for errors when closing a file, even in a deferred function.
func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

// Running the program confirms that the file is closed after being written