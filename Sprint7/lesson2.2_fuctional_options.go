package main

import "fmt"

// Object2 — объект с параметром.
type Object2 struct {
	// данные объекта
	// ...
	// настраиваемые поля объекта
	Mode int
	Path string
}

// WithMode — пример функции, которая присваивает поле Mode.
func WithMode(mode int) func(*Object2) {
	return func(o *Object2) {
		o.Mode = mode
	}
}

// WithPath — пример функции, которая присваивает поле Path.
func WithPath(path string) func(*Object2) {
	return func(o *Object2) {
		o.Path = path
	}
}

// NewObjectViaFuncOpts — функция-конструктор объекта.
func NewObjectViaFuncOpts(opts ...func(*Object2)) *Object2 {
	o := &Object2{}

	// вызываем все указанные функции для установки параметров
	for _, opt := range opts {
		opt(o)
	}
	return o
}

func main() {
	o := NewObjectViaFuncOpts(WithMode(10), WithPath(`root`))
	WithPath(`sad`)(o)
	fmt.Println(o)
}
