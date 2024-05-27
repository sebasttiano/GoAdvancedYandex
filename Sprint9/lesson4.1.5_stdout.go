package main

import (
	"fmt"
	"os"
)

const fname = "console.txt"

func main() {
	f, err := os.Create(fname)
	if err != nil {
		panic(err)
	}
	// сохраняем идентификатор текущего вывода
	stdout := os.Stdout
	// присваиваем os.Stdout идентификатор открытого файла
	os.Stdout = f
	// строка должна записаться в файл
	fmt.Println("Привет, гофер!")
	f.Close()

	// проверяем, что записалось в файл
	f, err = os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	out := make([]byte, 1024)
	var n int
	if n, err = f.Read(out); err != nil {
		panic(err)
	}
	// возвращаем обратно вывод в консоль
	// и смотрим, что прочитали из файла
	os.Stdout = stdout
	fmt.Println(string(out[:n]))
}
