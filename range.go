package main

import "fmt"

func main(){
// func ranges(){

// Range is used to iterate over arrays and slices

var sumArr, sumSlice int

// array
	var arr [5]int  = [5]int{20, 18, 7,42,72}
// calculate total of integers in arr

	for _,v := range arr {
	sumArr +=v
	}
	fmt.Println("Sum of array elements ")
	fmt.Println(sumArr)
 

// slice
	slice := []int{8,5,2,8,39,10, 35,683,373,36726,282,17}

	for _,v := range slice {
		sumSlice +=v
	 }
	 fmt.Println("Sum of slice elements ")
	 fmt.Println(sumSlice)
}