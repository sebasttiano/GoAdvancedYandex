package main

import (
	"fmt"
	"sync"
)

func main() {
	inCh := gen1(2, 3)
	ch1 := square2(inCh)
	ch2 := square2(inCh)
	for n := range fanIn1(ch1, ch2) {
		fmt.Println(n)
	}
}

func gen1(nums ...int) chan int {
	outCh := make(chan int)
	go func() {
		defer close(outCh)
		for _, n := range nums {
			outCh <- n
		}
	}()

	return outCh
}

func square2(inCh chan int) chan int {
	outCh := make(chan int)
	go func() {
		defer close(outCh)
		for n := range inCh {
			outCh <- n * n
		}
	}()

	return outCh
}

// fanIn1 принимает несколько каналов, в которых итоговые значения
func fanIn1(chs ...chan int) chan int {
	var wg sync.WaitGroup
	outCh := make(chan int)

	// определяем функцию output для каждого канала в chs
	// функция output копирует значения из канала с в канал outCh, пока с не будет закрыт
	output := func(c chan int) {
		for n := range c {
			outCh <- n
		}
		wg.Done()
	}

	// добавляем в группу столько горутин, сколько каналов пришло в fanIn1
	wg.Add(len(chs))
	// перебираем все каналы, которые пришли и отправляем каждый в отдельную горутину
	for _, c := range chs {
		go output(c)
	}

	// запускаем горутину для закрытия outCh после того, как все горутины отработают
	go func() {
		wg.Wait()
		close(outCh)
	}()

	// возвращаем общий канал
	return outCh
}
