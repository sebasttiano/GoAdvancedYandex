package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 1)

	// так как ёмкость канала больше 0, то можно записать
	// одно значение не ожидая, когда оно прочитается
	ch <- 10

	// попробуем закрыть канал, в котором есть значение
	close(ch)

	fmt.Println("len =", len(ch), "cap =", cap(ch))

	v, closed := <-ch
	fmt.Println(v, closed, "len =", len(ch))

	v, closed = <-ch
	fmt.Println(v, closed, "len =", len(ch))
}
