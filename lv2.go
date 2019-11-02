package main

import (
	"fmt"
	"sync"
)

var (
	myres = make(map[int]int, 20)
	mu    sync.Mutex
)
func factorial(n int) {
	var res = 1
	ch := make(chan int ,n)
	for i := 1; i <= n; i++ {
		res *= i
	}
	ch <- res
	myres[n] =<-ch

}

func main() {
	for i := 1; i <= 20; i++ {
		go factorial(i)
	}

	for i, a:= range myres {
		fmt.Printf("myres[%d] = %d\n", i, a)
	}

}
