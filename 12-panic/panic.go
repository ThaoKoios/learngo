// A panic typically means something went unexpecctedly wrong.
// Mostly, we use it to fail fast on errors that should not occur during normal opeartion,
// or that we are not prepared to handle gracefully.

package main

import "os"

func main() {

	// We will use panic throughout this site to check for unexpected errors.
	// This is the only program on the site designed to panic.
	panic("a problem")

	// Acommon use of panic is to abort if a function returns an error value that we don't know how to (or want to) handle.
	// Here's an example of `panicking` if we get an unexpected error when creating a new file
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}

// Running this program will cause it to panic, print an error message and goroutine traces, and exit with a non-zero status.
// Note that unlike some languages which use exceptions for handling of many errors,
// in Go it is idiomatic to use error-indication return values wherever possible
