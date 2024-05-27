package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	count := 3
	wg.Add(count)
	for i := 0; i < count; i++ {
		r := rand.New(rand.NewSource(int64(rnd.Uint64())))
		go func(i int) {
			fmt.Println(i, r.Float32(), r.Int31(), r.Intn(1000))
			wg.Done()
		}(i)
	}
	wg.Wait()
}
