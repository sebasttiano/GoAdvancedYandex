package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()                     // создаём новый список
	first := l.PushFront(2)             // добавляем одно число
	el := l.InsertBefore("zero", first) // вставляем перед ним другое число
	l.InsertAfter(1, el)                // вставляем 1 после предыдущего элемента

	// перебираем все элементы списка
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value)
	}
}
