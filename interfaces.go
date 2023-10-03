package main

import "fmt"

// interface

type Vehicle interface{
	Mileage() float64
	AverageSpeed() float64
}
// struct type 1
type Ford struct{
	distance float64
	fuel float64
	time float64
}

// struct type 2
type Toyota struct{
	distance float64
	fuel float64
	time float64
}

// implement methods
// Ford
func(vehicle Ford) Mileage() float64{
  return vehicle.distance/vehicle.fuel
}

func(vehicle Ford) AverageSpeed() float64{
  return vehicle.distance/vehicle.time
}

// Toyota
func(vehicle Toyota) Mileage() float64{
	return vehicle.distance/vehicle.fuel
  }
  
  func(vehicle Toyota) AverageSpeed() float64{
	return vehicle.distance/vehicle.time
  }


// func main(){
func interfaces(){

	// Ford
   ford := Ford{2000, 198, 51 }

   fmt.Println("ford.Mileage()")
   fmt.Println(ford.Mileage())
   fmt.Println("ford.AverageSpeed()")
   fmt.Println(ford.AverageSpeed())

   fmt.Println("=======")
   fmt.Println("=======")

   // Toyota
   toyota := Toyota{3690, 347, 92 }

   fmt.Println("toyota.Mileage()")
   fmt.Println(toyota.Mileage())
   fmt.Println("toyota.AverageSpeed()")
   fmt.Println(toyota.AverageSpeed())

}