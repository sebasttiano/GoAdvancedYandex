package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"time"
)

var (
	alphabet = []rune(`23456789@#$%^&/?+=!` +
		`ABCEFGHJKLMNPQRSTUVWXYZ` +
		`abcdefghijkmnopqrstuvwxyz`)
	alength = len(alphabet)
	rnd     *rand.Rand
	rnd2    *rand.Rand
)

// RandomString возвращает случайную шестнадцатеричную строку.
// Длина строки будет равна 2*n.
func RandomString(rnd *rand.Rand, n int) (string, error) {
	b := make([]byte, n)
	_, err := rnd.Read(b)
	if err != nil {
		return ``, err
	}
	return hex.EncodeToString(b), nil
}

// RandPsw возвращает случайный пароль указанной длины.
func RandPsw(count int) string {
	b := make([]rune, count)
	for i := range b {
		b[i] = alphabet[rnd2.Intn(alength)]
	}
	return string(b)
}

func main() {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	randMy := func(n int) string {
		s, err := RandomString(rnd, n)
		if err != nil {
			log.Fatal(err)
		}
		return s
	}
	fmt.Println(randMy(4), randMy(5), randMy(6))

	rnd2 = rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println(RandPsw(6), RandPsw(8), RandPsw(11))
}
