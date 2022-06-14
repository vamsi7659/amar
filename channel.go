package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println()
	mych := make(chan int, 2)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println(<-mych)
		//fmt.Println(<-mych)
		wg.Done()
	}(mych, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		mych <- 5
		mych <- 6
		wg.Done()
	}(mych, wg)
	wg.Wait()
}
