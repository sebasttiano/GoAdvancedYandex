//7. Ошибки — это значения (Errors are values).

package main

import "fmt"

type Storage map[string]string

func (s Storage) Get(key string) (string, error) {
	value, ok := s[key]
	if !ok {
		return "", fmt.Errorf("key not found")
	}
	return value, nil
}
