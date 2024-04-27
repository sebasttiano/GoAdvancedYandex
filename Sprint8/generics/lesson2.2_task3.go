package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

type Slice3[T any] []T

func (s *Slice3[T]) Map(f func(T) T) *Slice3[T] {
	for k, v := range *s {
		(*s)[k] = f(v)
	}
	return s
}

func (s *Slice3[T]) Reduce(r T, f func(a, e T) T) T {
	for _, v := range *s {
		r = f(r, v)
	}
	return r
}

func Sum2[T constraints.Ordered](a, e T) T {
	return a + e
}

func Double3[T constraints.Ordered](v T) T {
	return v + v
}

func main() {
	var si = Slice3[int]{1, 2, 3, 4, 5}
	sum := si.Reduce(0, Sum2[int])
	fmt.Println(sum)

	// теперь цепочку
	res := si.Map(Double3[int]).Reduce(0, Sum2[int])
	fmt.Println(res)

	// теперь для строк
	var ss = Slice3[string]{"foo", "bar", "buzz"}
	res1 := ss.Map(Double3[string]).Reduce("", Sum2[string])
	fmt.Println(res1)
}

func (s *Slice3[T]) Filter(allow func(e T) bool) *Slice3[T] {

	var res Slice3[T]
	for _, v := range *s {
		if allow(v) {
			res = append(res, v)
		}
	}
	return &res
}

func isEven(i int) bool {
	return i%2 == 0
}
