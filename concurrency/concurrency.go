package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("goroutine")
	fmt.Println("main")

	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println("gr:", i)
		}()
	}
	time.Sleep(10 * time.Millisecond)

	ch := make(chan string)
	go func() {
		ch <- "hi" // send
	}()
	val := <-ch // receive
	fmt.Println(val)

	fmt.Println(sleepSort([]int{30, 10, 20})) // [10 20 30]
}

/*
for every value "n" in values, spin a goroutine that
  - sleep n milliseconds
  - send n over a channel

collect all values from the channel to a slice and return it
*/
func sleepSort(values []int) []int {
	// FIXME
	// Hint: convert int to a float:
	// n := 1
	// f := float64(n)
	ch := make(chan int)
	for _, value := range values {
		value := value
		go func() {
			time.Sleep(time.Duration(value) * time.Millisecond)
			ch <- value
		}()
	}
	var sorted []int
	for range values {
		val := <-ch
		sorted = append(sorted, val)
	}
	return sorted
}

/* Channel semantics
- send/receive will block until opposite operation happen(*)
*/
