package main

import (
	"fmt"
	"time"
)

func writer(ch chan int) {
	for i := 0; ; i++ {
		time.Sleep(time.Second)
		ch <- i
	}
}
func reader(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	ch := make(chan int)
	go writer(ch)
	go reader(ch)

	timer := time.After(time.Second * time.Duration(n))
	<-timer
}
