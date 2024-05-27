package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rnd := rand.New(rand.NewSource(77))
	// всегда будет одна и та же последовательность
	// 460 3733561740 1284141128648234027
	fmt.Println(rnd.Intn(1000), rnd.Uint32(), rnd.Uint64())

	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
	// будут печататься разные числа
	fmt.Println(rnd.Intn(99), rnd.Uint32())
}
