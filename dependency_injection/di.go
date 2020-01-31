package main

import (
	"fmt"
	"io"
	"net/http"
)

// func Greet(writer *bytes.Buffer, name string) {
// 	fmt.Fprintf(writer, "Hello, %s", name)
// }
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "Arima")
}
func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}
