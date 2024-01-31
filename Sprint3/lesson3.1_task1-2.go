package main

import (
	"errors"
	"fmt"
	"os"
)

// LabelError описывает ошибку с дополнительной меткой.
type LabelError struct {
	Label string // метка должна быть в верхнем регистре
	Err   error
}

// Error добавляет поддержку интерфейса error для типа TimeError.
func (le *LabelError) Error() string {
	return fmt.Sprintf("[%v] %v", le.Label, le.Err)
}

// Unwrap возвращает исходную ошибку.
func (le *LabelError) Unwrap() error {
	return le.Err
}

// NewLabelError записывает ошибку err в тип TimeError c текущим временем.
func NewLabelError(label string, err error) error {
	return &LabelError{
		Label: label,
		Err:   err,
	}
}

// добавьте методы Error() и NewLabelError(label string, err error)

// ...

func main() {
	_, err := os.ReadFile("mytest.txt")
	if err != nil {
		err = NewLabelError("file", err)
	}
	fmt.Println(errors.Is(err, os.ErrNotExist), err)
	// должна выводить текст:
	// true [FILE] open mytest.txt: no such file or directory
}
