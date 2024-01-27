package main

import "net/http"

func mainPageMyMux(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Привет!"))
}

func apiPageMyMux(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Это страница /api."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc(`/api/`, apiPageMyMux)
	mux.HandleFunc(`/`, mainPageMyMux)
	// ...
	err := http.ListenAndServe(`:8080`, mux)
	// ...
	if err != nil {
		panic(err)
	}
}
