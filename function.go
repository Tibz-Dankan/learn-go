package main

import "fmt"

// Declaration
func sum(a int, b int) int {
	return a + b
}

// Function with multiple returns
func swap(a int, b int)(int, int) {
	return b, a
}

// Variadic function (function that any number of arguments)
func sumAll(nums ...int) int{
	var total int = 0
    
	for _, v := range nums{
		total +=v
	}
    return total
}

// Methods (functions with special receiver arguments)

type User struct{
	username string
	email string
	phone int
}


func (user *User) greet() string{
	greet := "Hello "+user.username
	return greet
}

func (user *User) changeUsername(name string) string{
	user.username = name
	return user.username
}






// func main(){
func functions(){

	// calling
	res := sum(3,7)
	fmt.Println(res)

	swapRes1, swapRes2 := swap(4,9)
	fmt.Println(swapRes1)
	fmt.Println(swapRes2)

	// calling variadic
	numbers := []int{1,2,3,4,5,6,7,8,9,10}
	resSumAll := sumAll(numbers...)
	fmt.Println("resSumAll")
	fmt.Println(resSumAll)

	// calling method
	user1 :=  User{username:"yourname1",email:"yourname1@gmail.com",phone: 756342189}
	

	greetRes := user1.greet()
	fmt.Println(greetRes)

	newUsername := user1.changeUsername("Dankan.T")
	fmt.Println(newUsername)


}