package go_ants

import (
	"fmt"
	"sync"
	"time"

	"github.com/panjf2000/ants"
)

func AntsFunc() {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Hello World!")
}

func AntsPool() {
	tmpPool, _ := ants.NewPool(50)

	runTimes := 1000
	var wg sync.WaitGroup
	syncCalculateSum := func() {
		AntsFunc()
		wg.Done()
	}
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = tmpPool.Submit(syncCalculateSum)
	}
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", tmpPool.Running())
	fmt.Printf("finish all tasks.\n")
}
