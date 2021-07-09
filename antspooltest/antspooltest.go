package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
	"github.com/panjf2000/ants/v2"
)

var sum int32

func myFunc(i interface{}) {
	n := i.(int32)
	atomic.AddInt32(&sum, n)
	fmt.Printf("run with %d\n", n)
}

func demoFunc() {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Hello World!")
}

func main() {
	defer ants.Release()

	runTimes := 100

	// Use the common pool.
	var wg sync.WaitGroup
	syncCalculateSum := func() {
		demoFunc()
		wg.Done()
	}
	var begin =time.Now()
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = ants.Submit(syncCalculateSum)
	}
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", ants.Running())
	fmt.Printf("finish all tasks.\n")
	var elapseTime = time.Now().Sub(begin)
	fmt.Println("耗时:", elapseTime)




	// Use the pool with a function,
	// set 10 to the capacity of goroutine pool and 1 second for expired duration.
	begin =time.Now()
	p, _ := ants.NewPoolWithFunc(100, func(i interface{}) {
		demoFunc()
		wg.Done()
	})
	defer p.Release()

	// Submit tasks one by one.
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = p.Invoke(int32(i))
	}
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", p.Running())
	fmt.Printf("finish all tasks, result is %d\n", sum)
	elapseTime = time.Now().Sub(begin)
	fmt.Println("耗时:", elapseTime)
}
