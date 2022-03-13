package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func RunCount2() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		count2("sheep")
		wg.Done()
	}()

	go func() {
		count2("fish")
		wg.Done()
	}()

	wg.Wait()
}

func count2(thing string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
