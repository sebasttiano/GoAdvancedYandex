package main

import "fmt"

// Object — объект с параметром.
type Object struct {
	// данные объекта
	// ...
	// настраиваемые поля объекта
	Mode int
	Path string
}

// SetMode — пример функции, которая присваивает поле Mode.
func (o *Object) SetMode(mode int) *Object {
	o.Mode = mode
	return o
}

// SetPath — пример функции, которая присваивает поле Path.
func (o *Object) SetPath(path string) *Object {
	o.Path = path
	return o
}

// NewObject — функция-конструктор объекта.
func NewObject() *Object {
	return &Object{}
}

func main() {
	o := NewObject().SetMode(10).SetPath(`root`)
	fmt.Println(o)
}
