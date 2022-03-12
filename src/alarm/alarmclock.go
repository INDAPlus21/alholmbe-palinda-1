package main

import (
	"fmt"
	"sync"
	"time"
)

func Remind(text string, delay time.Duration) {
	for {
		time.Sleep(delay)
		fmt.Printf("The time is %s: %s\n", time.Now().Format("15:04:05"), text)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		Remind("time to eat", 10*time.Second)
		wg.Done()
	}()
	go func() {
		Remind("time to eat", 30*time.Second)
		wg.Done()
	}()
	go func() {
		Remind("time to eat", 60*time.Second)
		wg.Done()
	}()

	wg.Wait()
}
