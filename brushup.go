package main

/*
An initial brush up on Golang:
  - Channel/buffered channel interactions, interfaces
  and soon to be added, HTML Client and syslogging
  examples with Go standard + some external libraries
*/

import (
  "math"
  "log/syslog"
  "fmt"
)

/*
Add the integer array elements together and
send into the channel input
*/
func testChannels(x []int, y chan int){
  sum := 0
  for _, v := range s {
    sum += v
  }
  c <- sum
}

// Multiple returns
func testMultipleRet() (int, int) {
	return 3, 7
}

func main() {
  z := []int{1,2,3,4,5,-7,-9,0,7,8,9,9}
  in := make(chan int) // Input channel
  go testChannels(z[:len(z)/2], in) // Send second half slice
  x := <-in
  fmt.Println(x)



	a, b := test1()
	fmt.Println(a)
	fmt.Println(b)

	_, c := test1()
	fmt.Println(c)
}

// For clojure
test123 := func(a, b int ) int { return a+b} (3, 4)


// structs, arch independent numeric types etc..
type Teststr struct {
  x,y int
  a,b,c float32
  A *[]int
  F func()
}

teststr1 := new(NewVarType)

x := make(chan int)       // Channel
y := make(chan int, 100)  // Buffered channels

type Test0 interface {    // Interface types
BlockSize() int
Enc(src, dst []byte)
Dec(src, dst []byte)
}
