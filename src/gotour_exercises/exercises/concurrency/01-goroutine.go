package concurrency

import (
	"fmt"
	"time"
)

func count1(thing string) {
	for i := 1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}

func RunCount1() {
	go count1("sheep")
	count1("fish")
}
