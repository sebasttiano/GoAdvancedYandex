package main

import (
	"crypto/aes"
	"crypto/rand"
	"fmt"
)

func generateRandom8(size int) ([]byte, error) {
	// генерируем криптостойкие случайные байты в b
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func main() {
	src := []byte("Слепой банкир") // данные, которые хотим зашифровать
	fmt.Printf("original: %s\n", src)

	// константа aes.BlockSize определяет размер блока, она равна 16 байтам
	key, err := generateRandom8(aes.BlockSize) // ключ шифрования
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	// получаем cipher.Block
	aesblock, err := aes.NewCipher(key)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	dst := make([]byte, aes.BlockSize) // зашифровываем
	aesblock.Encrypt(dst, src)
	fmt.Printf("encrypted: %x\n", dst)

	src2 := make([]byte, aes.BlockSize) // расшифровываем
	aesblock.Decrypt(src2, dst)
	fmt.Printf("decrypted: %s\n", src2)
}
