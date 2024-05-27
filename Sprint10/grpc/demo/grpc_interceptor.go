package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func clientInterceptor(ctx context.Context, method string, req interface{},
	reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption) error {
	// выполняем действия перед вызовом метода
	start := time.Now()

	// вызываем RPC-метод
	err := invoker(ctx, method, req, reply, cc, opts...)

	// выполняем действия после вызова метода
	if err != nil {
		log.Printf("[ERROR] %s,%v", method, err)
	} else {
		log.Printf("[INFO] %s,%v", method, time.Since(start))
	}
	return err
}

func main() {
	// устанавливаем соединение с сервером
	conn, err := grpc.Dial(":3200", grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(clientInterceptor))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
}
