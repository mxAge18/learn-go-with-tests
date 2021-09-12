package main

import (
	"fmt"
	"io"
	"net/http"
	// "os"
)

func Greet(buf io.Writer, name string) {
	fmt.Fprintf(buf, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "golang learner")
}
func main() {
	// Greet(os.Stdout, "maxu")
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}