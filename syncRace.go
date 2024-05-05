package goConcurrency

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	counter := 0

	gs := 100
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			defer wg.Done()
			for k := 0; k < 100; k++ {
				counter++
			}
		}()
	}
	wg.Wait()
	fmt.Println("Counter:", counter)
}
