package main

import (
	// ...

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"time"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer, TimerTrace)
	// или
	// r.Use(middleware.RealIP, middleware.Logger, middleware.Recoverer)

	// ...
}

func TimerTrace(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// перед началом выполнения функции сохраняем текущее время
		start := time.Now()
		// вызываем следующий обработчик
		next.ServeHTTP(w, r)
		// после завершения замеряем время выполнения запроса
		duration := time.Since(start)
		// сохраняем или сразу обрабатываем полученный результат
		// ...
		log.Printf("Функция выполнилась за: %d", duration)
	})
}
