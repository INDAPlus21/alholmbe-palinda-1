package library

import "fmt"

func fibonacci() func() int {
	i, j := 0, 1
	return func() int {
		k := i + j
		i = j
		j = k
		return k
	}
}

func RunFibo() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
