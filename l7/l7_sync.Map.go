package main

import (
	"fmt"
	"sync"
)

var cmap sync.Map

func Write(wg *sync.WaitGroup, i int) {
	defer wg.Done()
	cmap.Store(i, i+1)
}

func Read(wg *sync.WaitGroup, i int) {
	defer wg.Done()
	fmt.Println(cmap.Load(i))
}

func main() {
	wg := &sync.WaitGroup{}
	for i := range 10 {
		wg.Add(2)
		go Write(wg, i)
		go Read(wg, i)
	}
	wg.Wait()
}
