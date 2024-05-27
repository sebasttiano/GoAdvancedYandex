package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	// создаём файл для записи
	f, err := os.OpenFile("demo.txt", os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	var counter int
	// получаем текущее время в виде строки
	now := time.Now().Format("\r\n02.01.06 15:04:05")

	// чтобы сделать пример компактнее и не проверять err после каждого вызова,
	// продолжим выполнение после err == nil
	if finfo, err := f.Stat(); err == nil && finfo.Size() != 0 {
		// читаем текущий счётчик
		tmp := make([]byte, 4)
		// читаем текущее значение счётчика
		if _, err = f.ReadAt(tmp, 0); err == nil {
			counter, err = strconv.Atoi(string(tmp))
		}
	}
	if err != nil {
		log.Fatal(err)
	}
	// записываем значение счётчика в начало файла
	if _, err = f.WriteAt([]byte(fmt.Sprintf("%04d", (counter+1)%999)), 0); err == nil {
		// дописываем в конец текущее время
		if _, err = f.Seek(0, os.SEEK_END); err == nil {
			_, err = f.WriteString(now)
		}
	}
	if err != nil {
		log.Fatal(err)
	}
}
