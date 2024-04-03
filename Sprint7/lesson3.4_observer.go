package main

import (
	"context"
	"fmt"
	"time"
)

const shortDuration = 1 * time.Millisecond

type Subscriber struct {
	ID int
}

func (s Subscriber) Subscribe(ctx context.Context) {
	<-ctx.Done()
	fmt.Printf("Subscriber %v is notified\n", s.ID)
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), shortDuration)
	defer cancel()
	// создаём подписчиков
	s1 := Subscriber{1}
	go s1.Subscribe(ctx)
	s2 := Subscriber{2}
	go s2.Subscribe(ctx)
	s3 := Subscriber{3}
	go s3.Subscribe(ctx)
	time.Sleep(2 * time.Millisecond)
}
