package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range ch {
		time.Sleep(2 * time.Second)
		fmt.Printf("Worker %d: %d\n", id, data)
	}
}

func main() {
	wg := sync.WaitGroup{}
	defer wg.Wait()
	var n int
	ch := make(chan int)
	fmt.Print("Число: ")
	fmt.Scan(&n)
	if n < 2 {
		fmt.Println("error n<2")
		return
	}
	for i := range n {
		wg.Add(1)
		go worker(i, ch, &wg)
	}
	for i := 0; ; i++ {
		ch <- i
	}
}
