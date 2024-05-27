package main

import (
	"fmt"
	"math/rand"
	"time"
)

var rnd3 *rand.Rand

func main() {
	rnd3 = rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 100; i < 10001; i *= 10 {
		fmt.Println("n=", i, Prob(cube(i)))
	}
}

func cube(count int) []int8 {
	var raw []int8

	for i := 0; i < count; i++ {
		raw = append(raw, int8(rnd3.Int31n(6))+1)
	}
	return raw
}

func Prob(scores []int8) [6]float64 {
	var counts [6]int
	n := len(scores)
	for i := 0; i < n; i++ {
		counts[scores[i]-1]++
	}
	var av [6]float64
	for i := 0; i < 6; i++ {
		av[i] = float64(counts[i]) / float64(n)
	}
	return av
}
