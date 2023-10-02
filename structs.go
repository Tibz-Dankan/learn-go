package main

import "fmt"

// func main(){
func structs(){
	
// Declaration
type User struct{
	username string
	email string
	phone int
}

// create new struct

// including field names
user1 :=  User{username:"yourname1",email:"yourname1@gmail.com",phone: 756342189}
fmt.Println(user1)  //{yourname1 yourname1@gmail.com 756342189}

// without including field names
user2 :=  User{"yourname","yourname@gmail.com",756763215}
fmt.Println(user2)

// Accessing values in structs
username1 := user1.username
fmt.Println(username1)


// Pointers with structs
p :=  &User{"Dankan", "dankan@gmail.com", 789345671}

fmt.Println(p)  // &{Dankan dankan@gmail.com 789345671}
fmt.Println(p.username) //Struct  automatically dereferenced //Dankan
fmt.Println(p.email)  //Struct  automatically dereferenced   //dankan@gmail.com
fmt.Println(p.phone)   //Struct  automatically dereferenced  //789345671

}