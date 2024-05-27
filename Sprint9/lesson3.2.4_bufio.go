package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode/utf8"
)

const Lines = 1000

// GetTextFile создаёт тестовый текстовый файл.
func GetTextFile(fname string) (*os.File, error) {
	f, err := os.Create(fname)
	if err != nil {
		return nil, err
	}
	w := bufio.NewWriter(f)
	for i := 1; i <= Lines; i++ {
		fmt.Fprintf(w, "%d,Тестовое сообщение №%[1]d\r\n", i)
	}
	w.Flush()
	f.Seek(0, os.SEEK_SET)
	return f, nil
}

// ScanNumber находит очередное число во входящих данных.
func ScanNumber(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	start := 0
	// ищем первую цифру
	for width := 0; start < len(data); start += width {
		var r rune
		r, width = utf8.DecodeRune(data[start:])
		if r == utf8.RuneError { // ошибочный символ UTF-8
			if atEOF { // если конец, то возвращаем ошибку
				return 0, nil, fmt.Errorf("invalid rune")
			}
			return start, nil, nil // просим подгрузить данные
		}
		if r >= '0' && r <= '9' { // нашли цифру
			break
		}
	}
	for i := start + 1; i < len(data); i++ { // смотрим следующие символы
		if data[i] < '0' || data[i] > '9' { // закончилось число
			return i + 1, data[start:i], nil
		}
	}
	// входные данные заканчиваются числом
	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}
	// запрашиваем ещё данные, так как дальше тоже могут быть цифры
	return start, nil, nil
}

func main() {
	f, err := GetTextFile("my.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	// подключаем свой сканер чисел
	scanner.Split(ScanNumber)

	var (
		count  int64 // общее количество чисел
		count3 int64 // количество трёхзначных чисел
	)

	for scanner.Scan() {
		if len(scanner.Text()) == 3 {
			count3++
		}
		count++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(count, count3)
}
