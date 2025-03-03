package main

import (
	"fmt"
	"sync"
)

func task(id int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println("doing task", id)
}
/**

A WaitGroup in Go is a synchronization mechanism that waits for a collection of goroutines to finish execution before continuing the main function.
Think of it as a counter that:
Increases when a new goroutine starts (wg.Add(1)).
Decreases when a goroutine finishes (wg.Done()).
Blocks execution until all goroutines are done (wg.Wait())
*/

// Here we need to take care of wait groups as like , we need to add and then done and then delete 

func main() {
	var wg sync.WaitGroup
	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go task(i, &wg)
	}

	wg.Wait()
}
