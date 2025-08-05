package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func worker(sw *bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for *sw == false {
		time.Sleep(time.Second)
		fmt.Println("work")
	}
	fmt.Println("done")
}

func main() {
	stopWorker := false
	wg := &sync.WaitGroup{}
	defer wg.Wait()
	for _ = range 5 {
		wg.Add(1)
		go worker(&stopWorker, wg)
	}
	time.Sleep(time.Second * 5)
	stopWorker = true
}

//=================================================

package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func worker(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-time.After(time.Second):
			fmt.Println("work")
		case <-ch:
			fmt.Println("done")
			return
		}
	}
}

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan int)
	for _ = range 5 {
		wg.Add(1)
		go worker(ch, &wg)
	}
	time.Sleep(time.Second * 5)
	close(ch)
	wg.Wait()
}

// ====================================================
package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Done")
			return
		case <-time.After(time.Second):
			fmt.Println("Working...")
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	wg := &sync.WaitGroup{}
	for _ = range 5 {
		wg.Add(1)
		go worker(ctx, wg)
	}
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
}

// ========================================
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer fmt.Println("выполнение defer функции ")
	for i := range 5 {
		if i == id {
			fmt.Printf("runtime.Goexit у %d: ", id)
			runtime.Goexit()
		}
	}
}

func main() {
	wg := &sync.WaitGroup{}

	for i := range 5 {
		wg.Add(1)
		go worker(i, wg)
	}

	wg.Wait()
	fmt.Println("done")
}

// ===========================================
package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(stop bool, once *sync.Once, stopFunc func()) {
	once.Do(stopFunc)
	for !stop {
		time.Sleep(time.Second)
		fmt.Println("worrrk")
	}
	fmt.Println("done")
}
func main() {
	var once sync.Once // Сработает один раз
	stop := false

	stopper := func() {
		stop = true
		fmt.Println("Остановка началась")
	}

	go worker(stop, &once, stopper)
	time.Sleep(2 * time.Second)
	once.Do(stopper) // Попытка повторного вызова - не сработает
	time.Sleep(100 * time.Millisecond)
}
