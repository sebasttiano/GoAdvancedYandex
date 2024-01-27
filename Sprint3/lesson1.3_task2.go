package main

import (
	"context"
	"fmt"
	"time"
)

type DB2 struct {
}

type User2 struct {
	Name string
}

func (d *DB2) SelectUser(ctx context.Context, email string) (User2, error) {
	timer := time.NewTimer(1 * time.Second)
	select {
	case <-timer.C:
		return User2{Name: "Gosha"}, nil
	case <-ctx.Done():
		return User2{}, fmt.Errorf("context canceled")
	}
}

type Handler2 struct {
	db *DB2
}

type Request2 struct {
	Email string
}

type Response2 struct {
	User User2
}

func (h *Handler2) HandleAPI(ctx context.Context, req Request2) (Response2, error) {
	u, err := h.db.SelectUser(ctx, req.Email)
	if err != nil {
		return Response2{}, err
	}

	return Response2{User: u}, nil
}

func main() {
	db := DB2{}
	handler := Handler2{db: &db}

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	req := Request2{Email: "test@yandex.ru"}
	resp, err := handler.HandleAPI(ctx, req)
	fmt.Println(resp, err)
}
