package main

import "fmt"

// Map is used to create key value pairs
// Also known as dictionaries or associative arrays
// Note: Each key must be unique

// func main(){
func maps(){
	// Declaration of  a map
	m := make(map[string]int)

// Add more key-value pair to a map
  m["Dankan"]= 20
  m["Nathaniel"]= 16
  m["Elias"]= 8
  fmt.Println(m)  //map[Dankan:20 Elias:8 Nathaniel:16]

// Accessing values in the map
fmt.Println(m["Dankan"])
fmt.Println(m["Elias"])

// delete key value pair
fmt.Println()

delete(m, "Elias")

fmt.Println(m)

}