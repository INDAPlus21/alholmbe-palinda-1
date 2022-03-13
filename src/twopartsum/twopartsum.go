package main

import (
	"fmt"
)

// sum the numbers in a and send the result on res.
func sum(a []int, res chan<- int) {
	total := 0
	for _, value := range a {
		total += value
	}
	// send result on res
	res <- total
}

// concurrently sum the array a.
func ConcurrentSum(a []int) int {
	ch := make(chan int)

	go sum(a[:len(a)/2], ch)
	go sum(a[len(a)/2:], ch)

	// Get the subtotals from the channel and return their sum
	return <-ch + <-ch
}

func main() {
	// example call
	a := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(ConcurrentSum(a))
}
