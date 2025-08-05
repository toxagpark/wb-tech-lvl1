package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-time.After(time.Second * 1):
			fmt.Println("work")

		case <-ctx.Done():
			fmt.Println("ctx timeout")
			time.Sleep(time.Second * 5)
			fmt.Println("worker done")
			return
		}
	}
}

func main() {
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())
	chCancel := make(chan os.Signal, 1)
	signal.Notify(chCancel, syscall.SIGINT)

	for _ = range 5 {
		wg.Add(1)
		go worker(ctx, &wg)
	}
	<-chCancel
	cancel()
	wg.Wait()
}
