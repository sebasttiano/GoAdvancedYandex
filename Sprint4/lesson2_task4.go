package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		mu sync.Mutex
		wg sync.WaitGroup
	)
	m := make(map[int]int)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(v int) {
			for j := 0; j < 10; j++ {

				k := 10*v + j
				mu.Lock()
				m[k] = k
				mu.Unlock()
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(len(m))
}
