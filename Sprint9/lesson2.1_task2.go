package main

import (
	"crypto/rand"
	"fmt"
)

const Alphabet = 26

func GenCryptoKey(n int) (string, error) {

	rnd := make([]byte, n)
	nrnd, err := rand.Read(rnd)
	if err != nil {
		return ``, err
	} else if nrnd != n {
		return ``, fmt.Errorf(`nrnd %d != n %d`, nrnd, n)
	}
	for i := range rnd {
		rnd[i] = 'A' + rnd[i]%Alphabet
	}
	return string(rnd), nil

}

func main() {
	for i := 16; i <= 64; i += 16 {
		key, err := GenCryptoKey(i)
		if err != nil {
			panic(err)
		}
		fmt.Println(key)
	}
}
