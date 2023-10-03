//-> concurrency is running multiple processes independently
//-> Goroutines, channels and select are used to achieve concurrency
//-> Goroutines are virtual virtual threads that run independently
// of each and consume fewer resources compared to os threads
//-> channels are pipes that allow sending and receiving data from goroutines
//-> select is used when there is need to run some code after any of the
//  goroutines returns data via a channel

package main

import (
	"fmt"
	"time"
)

func evenSum(from, to int, ch chan int) {
    result := 0
    for i:=from; i<=to; i++ {
        if i%2 == 0 {
            result += i
        }    
    }
    ch <- result
}
func squareSum(from, to int, ch chan int) {
    result := 0
    for i:=from; i<=to; i++ {
        if i%2 == 0 {
            result += i*i
        }    
    }
    ch <- result
}

// func main() {
func concurrencies() {
    evenCh := make(chan int)
    sqCh := make(chan int)

    go evenSum(0, 100, evenCh)
    go squareSum(0, 100, sqCh)

    fmt.Println(<-evenCh + <- sqCh)

    select {
    case x := <- evenCh:
      fmt.Println(x)
    case y := <- sqCh:
        fmt.Println(y)  
    default :
        fmt.Println("No data returned any goroutine")  
        time.Sleep(50*time.Millisecond)
  }
}