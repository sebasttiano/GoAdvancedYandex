package main

import (
	"bytes"
	"fmt"
	"log"
)

func main() {
	var buf bytes.Buffer
	// допишите код
	// 1) создайте переменную типа *log.Logger
	// 2) запишите в неё нужные строки
	myLog := log.New(&buf, "mylog: ", 0)
	// ...
	myLog.Println("Hello, world")
	myLog.Println("Goodbye")
	fmt.Print(&buf)
	// должна вывести
	// mylog: Hello, world!
	// mylog: Goodbye
}
