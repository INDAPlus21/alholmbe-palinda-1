package concurrency

import (
	"fmt"
	"strconv"
	"time"
)

func RunCount3() {
	c := make(chan string, 50)
	go count3("sheep", c)

	for msg := range c {
		fmt.Println(msg)
	}
}

func count3(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing + strconv.FormatInt(int64(i), 10)
		time.Sleep(time.Millisecond * 500)
	}

	close(c)
}
