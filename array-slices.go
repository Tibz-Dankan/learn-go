package main

import "fmt"

// func main(){
func arraySlice(){

	// Array
	// Declaration 

	var ages [5]int  = [5]int{20, 18, 7,42,72}
	fmt.Println(ages)  //[20 18 7 42 72]

	age  := [5]int{23, 34,98,61,3}
	fmt.Println(age)   //[23 34 98 61 3]

	// Accessing array elements

	var element1 int = age[2]
	fmt.Println(element1)  //98

	element2 := ages[3]
	fmt.Println(element2)  //42


	// Slices - Provide dynamic sized view into array elements
	// Declaration
	s1 := age[1:4] //default indices- low is 0 and high is array length
	fmt.Println(s1)  //[34 98 61]

	s2 := []int{8,5,2,8,39,10, 35,683,373,36726,282,17}
	fmt.Println(s2) //[8 5 2 8 39 10 35 683 373 36726 282 17]

	// Modify array element
	s1[0]= 69
	fmt.Println(age)  //[23 69 98 61 3]

	// create a slice
	newSlice := make([]int,5)
	fmt.Println(newSlice)  //[0 0 0 0 0]
	// Add elements to slice
	newAges := append(newSlice, 1,3,4,7,26)
	fmt.Println(newAges)  //[0 0 0 0 0 1 3 4 7 26]

}