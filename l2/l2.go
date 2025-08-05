package main

import (
	"fmt"
	"sync"
)

func makeSquare(a int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(a * a)
}

func main() {
	wg := sync.WaitGroup{}
	s := []int{2, 4, 6, 8, 10}
	for _, i := range s {
		wg.Add(1)
		go makeSquare(i, &wg)
	}
	wg.Wait()
}
