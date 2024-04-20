package main

import (
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

func main() {
	// создаём файл журнала профилирования cpu
	fcpu, err := os.Create(`cpu.profile`)
	if err != nil {
		panic(err)
	}
	defer fcpu.Close()
	if err := pprof.StartCPUProfile(fcpu); err != nil {
		panic(err)
	}
	defer pprof.StopCPUProfile()

	go foo()
	time.Sleep(10 * time.Second)

	// создаём файл журнала профилирования памяти
	fmem, err := os.Create(`mem.profile`)
	if err != nil {
		panic(err)
	}
	defer fmem.Close()
	runtime.GC() // получаем статистику по использованию памяти
	if err := pprof.WriteHeapProfile(fmem); err != nil {
		panic(err)
	}
}
