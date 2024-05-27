package main

import (
	"golang.org/x/time/rate"
	"net/http"
)

// limiter разрешает 3 запроса в секунду.
var limiter = rate.NewLimiter(3, 1)

// CheckTimeRate — обработчик middleware для проверки ограничений.
func CheckTimeRate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, `Слишком много запросов`, http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Привет!"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", DefaultHandler)
	http.ListenAndServe(":3000", CheckTimeRate(mux))
}
