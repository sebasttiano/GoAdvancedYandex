package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

type Iterator[V any] interface {
	First() (V, bool)
	Next() (V, bool)
	Set(V)
}

func Map[I Iterator[V], V any](it I, f func(V) V) {
	for item, ok := it.First(); ok; item, ok = it.Next() {
		it.Set(f(item))
	}
}

// Item — элемент списка со ссылкой на следующий элемент.
type Item[T any] struct {
	next  *Item[T]
	value T
}

// List — список.
type List[T any] struct {
	first *Item[T] // первый элемент
	cur   *Item[T] // текущий элемент
}

func NewList[V any](l *Item[V]) *List[V] {
	var i List[V]
	i.first = l
	i.cur = l
	return &i
}

func (l *List[T]) Next() (T, bool) {
	if l.cur.next != nil {
		l.cur = l.cur.next
		return l.cur.value, true
	}
	var empty T
	return empty, false
}

func (l *List[T]) First() (T, bool) {
	l.cur = l.first
	if l.cur == nil {
		var empty T
		return empty, false
	}
	return l.cur.value, true
}

func (l *List[T]) Set(v T) {
	if l.cur != nil {
		l.cur.value = v
	}
}

func Double4[T constraints.Ordered](v T) T {
	return v + v
}

func main() {
	// создаём первый элемент
	li := new(Item[int])
	// добавляем элементы
	for i := 1; i < 7; i++ {
		nl := new(Item[int])
		nl.value = i
		nl.next = li
		li = nl
	}
	// конструируем список
	var list = NewList(li)
	for item, ok := list.First(); ok; item, ok = list.Next() {
		fmt.Println(item)
	}
	Map(list, Double4[int])
	fmt.Println("After Mapping")
	for item, ok := list.First(); ok; item, ok = list.Next() {
		fmt.Println(item)
	}

	m := map[string]string{
		"1": "a",
		"2": "b",
		"3": "c",
		"4": "d",
		"5": "e",
	}
	// конструируем итератор для мапы
	var iter = NewMapIter(m)

	for item, ok := iter.First(); ok; item, ok = iter.Next() {
		fmt.Println(item)
	}
	// применяем метод Map()
	Map(iter, Double4[string])
	fmt.Println("After Mapping")

	for item, ok := iter.First(); ok; item, ok = iter.Next() {
		fmt.Println(item)
	}
}

type MapGeneric[K constraints.Ordered, V any] map[K]V

type MapIter[K constraints.Ordered, V any] struct {
	m     MapGeneric[K, V]
	index []K
	cur   int
}

func NewMapIter[K constraints.Ordered, V any](m MapGeneric[K, V]) *MapIter[K, V] {
	var ret MapIter[K, V]
	ret.m = m
	ret.index = make([]K, 0, len(m))
	for k := range m {
		ret.index = append(ret.index, k)
	}
	return &ret
}

func (m *MapIter[K, V]) Next() (V, bool) {
	if m.cur < len(m.index)-1 {
		m.cur++
		return m.m[m.index[m.cur]], true
	}
	var empty V
	return empty, false
}

func (m *MapIter[K, V]) Set(v V) {
	if m.cur < len(m.index) {
		m.m[m.index[m.cur]] = v
	}
}

func (m *MapIter[K, V]) First() (V, bool) {
	m.cur = 0
	if len(m.index) > 0 {
		return m.m[m.index[m.cur]], true
	}
	var empty V
	return empty, false
}
