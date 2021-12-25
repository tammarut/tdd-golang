package io_example_test

import (
	"errors"
	"io"
	"reflect"
	"testing"

	"github.com/tammarut/tdd-golang/mocking/io_example"
)

type mockReadCloser struct {
	expectedContents []byte
	expectedError    error
}

func (mockReadCloser *mockReadCloser) Read(p []byte) (n int, err error) {
	copy(p, mockReadCloser.expectedContents)
	return 0, mockReadCloser.expectedError
}

func (mockReadCloser *mockReadCloser) Close() error {
	return nil
}

type testCase struct {
	name             string
	readCloser       io.ReadCloser
	numBytes         int
	expectedContents []byte
	expectedError    error
}

func TestReadContents(t *testing.T) {
	// Arrange
	errRead := errors.New("Read file contents error")
	testCases := []testCase{
		{
			name: "Happy case",
			readCloser: &mockReadCloser{
				expectedContents: []byte(`hello`),
				expectedError:    nil,
			},
			numBytes:         5,
			expectedContents: []byte(`hello`),
			expectedError:    nil,
		},
		{
			name: "Failed case",
			readCloser: &mockReadCloser{
				expectedContents: nil,
				expectedError:    errRead,
			},
			numBytes:         0,
			expectedContents: nil,
			expectedError:    errRead,
		},
	}

	// Act
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			contents, err := io_example.ReadFileContents(testCase.readCloser, testCase.numBytes)

			// Assert
			if !reflect.DeepEqual(contents, testCase.expectedContents) {
				t.Errorf("expected %b, but got %b", testCase.expectedContents, contents)
			}
			if !errors.Is(err, testCase.expectedError) {
				t.Errorf("expected %v, but got error %v", testCase.expectedError, err)

			}
		})

	}

}
