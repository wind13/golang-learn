package main

import "fmt"
import "golang.org/x/tour/reader"

type MyReader struct{}

type ErrRead int

func (e ErrRead) Error() string {
	// return fmt.Sprintf("cannot Sqrt negative number: %v\n", float64(e))
	return fmt.Sprintf("cannot read number: %d\n", e)
}

// Add a Read([]byte) (int, error) method to MyReader.
func (reader MyReader) Read([]byte) (int, error) {
	var b = make([]byte)
	return 0, ErrRead(b)
}

func main() {
	reader.Validate(MyReader{})
}
