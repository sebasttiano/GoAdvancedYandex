package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// получаем все пути из переменной PATH
	paths := strings.Split(os.Getenv("PATH"), string(os.PathListSeparator))
	// оставляем только три первых
	if len(paths) > 3 {
		paths = paths[:3]
	}
	// удаляем текущие и устанавливаем новые переменные окружения
	os.Clearenv()
	os.Setenv("PATH", strings.Join(paths, string(os.PathListSeparator)))
	os.Setenv("MYPATH", os.Args[0])
	for _, item := range os.Environ() {
		fmt.Println(item)
	}
	os.Setenv("MYAPP", filepath.Base(os.Args[0]))
	fmt.Println(os.ExpandEnv("PATH=$PATH\r\nAPP=${MYAPP}"))
}
