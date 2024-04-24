package main

import (
	"fmt"
	"sync"
	"time"
)

func abc(a string, ch chan int, wg *sync.WaitGroup) {
	fmt.Println(a)
	wg.Done()
	ch <- 10
}
func xyz(b string, ch chan int, wg *sync.WaitGroup) {
	fmt.Println(b)
	wg.Done()
	ch <- 20
}
func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go abc("Function one", ch, &wg)
	//time.Sleep(time.Second)
	fmt.Println(<-ch)
	go xyz("Function two", ch, &wg)
	time.Sleep(time.Second)
	fmt.Println(<-ch)
	wg.Wait()
}
