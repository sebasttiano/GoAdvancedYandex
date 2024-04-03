package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func ConditionalWrite(w io.Writer) (int, error) {
	switch val := w.(type) {
	case *os.File:
		fmt.Printf("%v\n", val)
		return 1, nil
	}
	// используем рефлексию для определения конкретного типа
	if wtype := reflect.TypeOf(w); wtype.String() == "*os.File" {
		// если это *os.File, предпринимаем действия
	}
	fmt.Printf("fff\n")
	return 1, nil
	// действуем по-другому
}
