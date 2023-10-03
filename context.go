// HTTP servers are useful for demonstrating the usage of
//  context.Context for controlling cancellation. A Context
//  carries deadlines, cancellation signals, and other
//  request-scoped values across API boundaries and goroutines.

package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello2(w http.ResponseWriter, req *http.Request) {

    ctx := req.Context()
    fmt.Println("server: hello handler started")
    defer fmt.Println("server: hello handler ended")

    select {
    case <-time.After(10 * time.Second):
        fmt.Fprintf(w, "hello\n")
    case <-ctx.Done():

        err := ctx.Err()
        fmt.Println("server:", err)
        internalError := http.StatusInternalServerError
        http.Error(w, err.Error(), internalError)
    }
}

// func main() {
func contexts() {

    http.HandleFunc("/hello", hello2)
    http.ListenAndServe(":8090", nil)
}