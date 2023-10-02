package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func helloWord(w http.ResponseWriter, req *http.Request) {

    fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func clientHandler(w http.ResponseWriter, req *http.Request){
//     func (*http.Request).FormFile(key string) (multipart.File, *multipart.FileHeader, error)
// func (*http.Request).FormValue(key string) string
    fmt.Println(req.Body)
    fmt.Println(req.FormFile("file"))
    fmt.Println(req.FormFile("files"))
}

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

func person(w http.ResponseWriter, r *http.Request) {
    // Create a sample Person struct
    person := Person{
        FirstName: "John",
        LastName:  "Doe",
        Age:       30,
    }

    // Marshal the Person struct into JSON
    jsonData, err := json.Marshal(person)
    if err != nil {
        http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
        return
    }

    // Set the Content-Type header to indicate JSON
    w.Header().Set("Content-Type", "application/json")

    // Write the JSON response
    _, err = w.Write(jsonData)
    if err != nil {
        http.Error(w, "Unable to write response", http.StatusInternalServerError)
        return
    }
}

// func main() {
func httpServer() {

    http.HandleFunc("/hello", helloWord)
    http.HandleFunc("/headers", headers)
    http.HandleFunc("/client", headers)
    http.HandleFunc("/person", headers)

    http.ListenAndServe(":8090", nil)
}