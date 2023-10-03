// Bench to compared  time to taken run a program using concurrency v sequential

// Loops up to  1,000 iterations time is same for both concurrency and sequential
// Loops up to  1000,000 iterations concurrency is twice as fast as  sequential
// -: Concurrency is a winner

package main

import (
	"fmt"
	"time"
)

// with concurrency
func evenSums(from, to int, ch chan int) {
    result := 0
    for i:=from; i<=to; i++ {
        if i%2 == 0 {
            result += i
        }    
    }
    ch <- result
}

func squareSums(from, to int, ch chan int) {
    result := 0
    for i:=from; i<=to; i++ {
        if i%2 == 0 {
            result += i*i
        }    
    }
    ch <- result
}

// without concurrency
func evenSumy(from, to int) int {
    result := 0
    for i:=from; i<=to; i++ {
        if i%2 == 0 {
            result += i
        }    
    }
  return result
}

func squareSumy(from, to int) int {
    result := 0
    for i:=from; i<=to; i++ {
        if i%2 == 0 {
            result += i*i
        }    
    }
    return result
}

// func main() {
func benchMark(){

	start := time.Now()

	// ==> concurrency
    evenCh := make(chan int)
    sqCh := make(chan int)

    go evenSums(0, 1000000000, evenCh)
    go squareSums(0, 1000000000, sqCh)

    fmt.Println(<-evenCh + <- sqCh)

	// // ==> without concurrency
	// evenResult := evenSumy(0,1000000000)
	// sqResult := squareSumy(0,1000000000)
	// fmt.Println(evenResult + sqResult)

	end := time.Now()
    elapsed := end.Sub(start)
    elapsedMilliseconds := elapsed.Milliseconds()
    fmt.Printf("Function took %d milliseconds to run\n", elapsedMilliseconds)
 

   
}