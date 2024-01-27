package main

import (
	"log"
	"net/http"
)

// middleware принимает параметром Handler и возвращает тоже Handler.
func middleware(next http.Handler) http.Handler {
	// получаем Handler приведением типа http.HandlerFunc
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// здесь пишем логику обработки
		// например, разрешаем запросы cross-domain
		// w.Header().Set("Access-Control-Allow-Origin", "*")
		// ...
		// замыкание: используем ServeHTTP следующего хендлера
		next.ServeHTTP(w, r)
	})
}

func rootHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Привет"))
}

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://yandex.ru/", http.StatusMovedPermanently)
}

type Middleware func(http.Handler) http.Handler

// Если функций-обёрток много, можно подключить их с помощью вспомогательной функции.

func Conveyor(h http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		h = middleware(h)
	}
	return h
}

func main() {
	http.Handle("/", middleware(http.HandlerFunc(rootHandle)))
	http.HandleFunc("/search", redirect)
	log.Fatal(http.ListenAndServe(":8080", nil))
	//...
}

//func main() {
//	http.Handle("/", Conveyor(http.HandlerFunc(rootHandle), middleware1, middleware2, middleware3))
//	// ...
//}
