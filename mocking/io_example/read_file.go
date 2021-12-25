package io_example

import (
	"io"
)

/*
➜ I need to mock a method on a type (struct/class)
  - "Accept interfaces, return structs"
  - We used io.ReadCloser(built-in interface), but don't be afraid to define your own interface
*/
func ReadFileContents(readCloser io.ReadCloser, numBytes int) ([]byte, error) {
	defer readCloser.Close()
	contents := make([]byte, numBytes)
	_, err := readCloser.Read(contents)
	if err != nil {
		return nil, err
	}
	return contents, nil
}
