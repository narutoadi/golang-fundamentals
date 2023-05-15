package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Microsecond)
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 50; i++ {
		wg.Add(1)

		i := i

		// Wrap the worker call in a closure that makes sure to tell the WaitGroup that this worker is done.
		// This way the worker itself does not have to be aware of the concurrency primitives involved in its execution.
		go func() {
			defer wg.Done()
			worker(i)

		}()
	}

	wg.Wait()

}

// Note that this approach has no straightforward way to propagate errors from workers.
// For more advanced use cases, consider using the errgroup package.
