package main

import (
	"expvar"
	"io"
	"net/http"
	"time"
)

var (
	start = time.Now()
	// регистрируем счётчик
	numberOfRequests = expvar.NewInt("system.numberOfRequest")
)

// Hello обрабатывает HTTP-запрос.
func Hello(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Hello")
	numberOfRequests.Add(1) // увеличиваем счётчик запросов
}

// Uptime возвращает количество секунд с момента старта программы.
func Uptime() interface{} {
	return int64(time.Since(start).Seconds())
}

func main() {
	// регистрируем метрику uptime с вызовом функции Uptime
	expvar.Publish("system.uptime", expvar.Func(Uptime))

	http.HandleFunc("/", Hello)
	http.ListenAndServe(":1234", nil)
}
