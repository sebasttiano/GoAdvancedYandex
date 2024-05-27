package main

import (
	"fmt"
	"os"
)

func main() {
	// проверяем значение переменной окружения
	if env := os.Getenv("MYAPP"); len(env) > 0 {
		fmt.Printf("MYAPP=%s\r\n", env)
		return
	}
	// получаем имя приложения
	name, err := os.Executable()
	if err != nil {
		panic(err)
	}
	var procAttr os.ProcAttr
	// передаём при запуске переменную окружения MYAPP
	procAttr.Env = []string{"MYAPP=" + name}
	procAttr.Files = []*os.File{os.Stdin, os.Stdout, os.Stderr}
	// запускаем приложение
	proc, err := os.StartProcess(name, []string{name}, &procAttr)
	if err != nil {
		panic(err)
	}
	// ждём окончания работы запущенного приложения
	state, err := proc.Wait()
	if err != nil {
		panic(err)
	}
	fmt.Println("ExitCode", state.ExitCode())
}
