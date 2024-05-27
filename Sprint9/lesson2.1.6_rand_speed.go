package main

import (
	"math/rand"
	"testing"
	"time"

	crand "crypto/rand"
)

func BenchmarkMathRand(b *testing.B) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	data := make([]byte, 256)
	for i := 0; i < b.N; i++ {
		_, err := rnd.Read(data)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkCryptoRand(b *testing.B) {
	data := make([]byte, 256)
	for i := 0; i < b.N; i++ {
		_, err := crand.Read(data)
		if err != nil {
			panic(err)
		}
	}
}
