package main

import (
	"fmt"
)

func writeX(ch chan<- int, x [7]int) { //не использую wg.Done, так как горутина выполнится только тогда,
	// когда запишет все числа из массива, при этом ждет читателя при каждой записи
	defer close(ch)
	for _, i := range x {
		ch <- i
	}
}

func writeXSq(inputCh <-chan int, outputCh chan<- int) { //читает числа, возводит в квадрат, ждет писателя, заканчивает тогда,
	//когда закрывается канал в первой горутине
	defer close(outputCh)
	for i := range inputCh {
		x := i
		x *= x
		outputCh <- x
	}
}

func startWork(inputCh, outputCh chan int, x [7]int) { //горутины обязательно запускать вместе
	go writeX(inputCh, x)
	go writeXSq(inputCh, outputCh)
	for i := range outputCh {
		fmt.Print(i, " ")
	}
}

func main() {
	inputCh := make(chan int)
	outputCh := make(chan int)

	x := [7]int{1, 2, 5, 10, 15, 25, -10}

	startWork(inputCh, outputCh, x)
}
