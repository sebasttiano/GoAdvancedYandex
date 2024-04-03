package main

import (
	"fmt"
	"sync"
	"time"
)

type Subscriber2 struct {
	ID int
}

// Subscribe ожидает уведомления.
func (s Subscriber2) Subscribe(c *sync.Cond) {
	for {
		c.L.Lock()
		c.Wait()
		fmt.Printf("Subscriber2 %v is notified\n", s.ID)
		c.L.Unlock()
	}
}

func main() {
	cond := sync.NewCond(new(sync.Mutex))
	s1 := Subscriber2{1}
	go s1.Subscribe(cond)
	s2 := Subscriber2{2}
	go s2.Subscribe(cond)
	s3 := Subscriber2{3}
	go s3.Subscribe(cond)
	time.Sleep(1 * time.Second)
	// отправка уведомлений всем подписчикам
	cond.Broadcast()
	time.Sleep(1 * time.Second)
	fmt.Println("Once more")
	cond.Broadcast()
	time.Sleep(1 * time.Second)
}
