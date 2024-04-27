package main

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
)

func Foo(m sync.Mutex) {
	m.Lock()
	defer m.Unlock()
	// какие-то действия
}

type Foo2 struct {
	Name string `json: "nickname"`
}

func CtxFunc(ctx context.Context) {}

func main() {
	ctx, _ := context.WithCancel(context.Background())
	CtxFunc(ctx)
	j, err := json.Marshal(Foo2{Name: "JohnDoe"})
	fmt.Println(string(j), err)
	var i int
	if true {
		i := 7
		fmt.Print(i)
	}
	fmt.Print(i)
	var wg sync.WaitGroup
	for _, v := range []int{0, 1, 2, 3} {
		wg.Add(1)
		go func() {
			fmt.Print(v)
			wg.Done()
		}()
	}
	wg.Wait()
}
