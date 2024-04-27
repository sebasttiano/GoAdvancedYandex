package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

type Slice2[T any] []T

func (s *Slice2[T]) Map(f func(T) T) *Slice2[T] {
	for k, v := range *s {
		(*s)[k] = f(v)
	}
	return s
}

func (s *Slice2[T]) Reduce(r T, f func(a, e T) T) T {
	for _, v := range *s {
		r = f(r, v)
	}
	return r
}

func Sum[T constraints.Ordered](a, e T) T {
	return a + e
}

func Double2[T constraints.Ordered](v T) T {
	return v + v
}

func main() {
	var si = Slice2[int]{1, 2, 3, 4, 5}
	sum := si.Reduce(0, Sum[int])
	fmt.Println(sum)

	// теперь цепочку
	res := si.Map(Double2[int]).Reduce(0, Sum[int])
	fmt.Println(res)

	// теперь для строк
	var ss = Slice2[string]{"foo", "bar", "buzz"}
	res1 := ss.Map(Double2[string]).Reduce("", Sum[string])
	fmt.Println(res1)
}
