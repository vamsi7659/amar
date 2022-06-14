package main

import (
	"fmt"
	"sync"
)

func main() {
	scores := []int{0}
	//it does not allow goroutines to read or write while its running by using lock and unlock functions
	mut := &sync.RWMutex{}
	wg := &sync.WaitGroup{}
	wg.Add(3)
	go func(wg *sync.WaitGroup, m*sync.RWMutex) {

		fmt.Println("one")
		mut.Lock()
		scores =append(scores, 1)
		mut.Unlock()
		wg.Done()
	}(wg,mut)
	go func(wg *sync.WaitGroup,m*sync.RWMutex) {
		fmt.Println("two")
		mut.Lock()
		scores =append(scores, 2)
		mut.Unlock()
		wg.Done()
	}(wg,mut)
	go func(wg *sync.WaitGroup,m*sync.RWMutex) {
		fmt.Println("three")
		mut.Lock()
		scores =append(scores, 3)
		mut.Unlock()
		wg.Done()
	}(wg,mut)
	wg.Wait()
	fmt.Println(scores)
}