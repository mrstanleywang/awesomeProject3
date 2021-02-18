package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("begin CPU"runtime.NumCPU()N)
	fmt.Println("begin gs"runtime.NumGoroutine()N)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("hello beautiful real world")
		wg.Done()
	}()

	go func() {
		fmt.Println("hello beautiful ideal world")
		wg.Done()
	}()


	wg.Wait()
	fmt.Println("about to exit")
}
