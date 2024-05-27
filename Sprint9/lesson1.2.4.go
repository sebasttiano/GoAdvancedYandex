package main

import (
	"container/list"
	"sync"
)

type ListChannel struct {
	Lst *list.List
	Len int
	sync.Mutex
}

func MakeListChannel(l int) ListChannel {
	return ListChannel{Lst: list.New(), Len: l}
}

func (lc ListChannel) Send(v any) {
	for {
		lc.Lock()
		if lc.Lst.Len() < lc.Len {
			lc.Lst.PushBack(v)
			lc.Unlock()
			return
		}
		lc.Unlock()
	}
}

func (lc ListChannel) Receive() any {
	for {
		lc.Lock()
		if lc.Lst.Len() > 0 {
			e := lc.Lst.Front()
			lc.Lst.Remove(e)
			lc.Unlock()
			return e.Value
		}
		lc.Unlock()
	}
}
