package main

import (
	"fmt"
	"sort"
)

func main() {
	ss := sort.StringSlice{"foo", "bar", "buzz"}
	ss.Sort()
	fmt.Println(ss)
	n := ss.Search("foo") // бинарный поиск в отсортированном массиве
	fmt.Println(n)
}
