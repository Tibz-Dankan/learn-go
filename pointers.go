package main

import "fmt"

// Pointer is variable that holds memory address of another variable

// func main(){
func pointers(){

// Declaration
var p *int

age := 22

// Assignment
p = &age

// Modify value
*p = 23

fmt.Println(p) //Memory address // 0xc00009e010
fmt.Println(*p) // Actual using dereferencing  // 23

}