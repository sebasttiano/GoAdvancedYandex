package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

var secretkey = []byte("secret key")

func main() {
	var (
		data []byte // декодированное сообщение с подписью
		id   uint32 // значение идентификатора
		err  error
		sign []byte // HMAC-подпись от идентификатора
	)
	msg := "048ff4ea240a9fdeac8f1422733e9f3b8b0291c969652225e25c5f0f9f8da654139c9e21"

	data, err = hex.DecodeString(msg)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
	id = binary.BigEndian.Uint32(data[:4])
	fmt.Println(id)

	h := hmac.New(sha256.New, secretkey)
	fmt.Println(h)
	h.Write(data[:4])
	fmt.Println(h)
	sign = h.Sum(nil)
	fmt.Println(sign)
	// допишите код
	// 1) декодируйте msg в data
	// 2) получите идентификатор из первых четырёх байт,
	//    используйте функцию binary.BigEndian.Uint32
	// 3) вычислите HMAC-подпись sign для этих четырёх байт

	// ...

	if hmac.Equal(sign, data[4:]) {
		fmt.Println("Подпись подлинная. ID:", id)
	} else {
		fmt.Println("Подпись неверна. Где-то ошибка")
	}
}
