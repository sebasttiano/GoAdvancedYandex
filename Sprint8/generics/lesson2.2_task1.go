package main

import (
	"fmt"
	"testing"
)

// реализуйте функцию Reverse
// ...
func main() {
	fmt.Println(Reverse([]int{1, 2, 3, 4, 5}))
}

func Reverse[T int | string | float64](s []T) []T {

	revSlice := make([]T, len(s))
	for i, v := range s {
		revSlice[len(s)-i-1] = v
	}
	return revSlice
}

func TestReverse(t *testing.T) {
	if fmt.Sprint(Reverse([]int{10, -6, 34, 54})) != "[54 34 -6 10]" {
		t.Errorf(`wrong []int reverse`)
	}
	if fmt.Sprint(Reverse([]string{"foo", "buzz", "generic", "go"})) != "[go generic buzz foo]" {
		t.Errorf(`wrong []string reverse`)
	}
	if fmt.Sprint(Reverse([]float64{4.67, 5, -2.34, 7.88, 100})) != "[100 7.88 -2.34 5 4.67]" {
		t.Errorf(`wrong []float64 reverse`)
	}
}
