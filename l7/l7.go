package main

import (
	"fmt"
	"sync"
)

type SafeOperation struct {
	m       sync.Mutex
	safeMap map[int]int
}

func NewSafeOperation() *SafeOperation {
	return &SafeOperation{
		safeMap: make(map[int]int),
	}
}

func (so *SafeOperation) Write(a int, wg *sync.WaitGroup) {
	defer wg.Done()
	so.m.Lock()
	defer so.m.Unlock()
	so.safeMap[a] = a + 1
}
func (so *SafeOperation) Read(a int, wg *sync.WaitGroup) {
	defer wg.Done()
	so.m.Lock()
	defer so.m.Unlock()
	if b, ok := so.safeMap[a]; ok {
		fmt.Println(b)
	} else {
		fmt.Println(-1)
	}
}

func main() {
	wg := &sync.WaitGroup{}
	so := NewSafeOperation()
	for i := range 10 {
		wg.Add(2)
		go so.Write(i, wg)
		go so.Read(i, wg)
	}
	wg.Wait()
}
