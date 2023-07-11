package main

import (
	"fmt"
	"sync"
)

func main() {
	stock := 1000
	mutex := sync.Mutex{}
	group := sync.WaitGroup{}
	group.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			mutex.Lock()
			stock -= 1
			mutex.Unlock()
			group.Done()
		}()
	}
	group.Wait()
	fmt.Println(stock)
}
