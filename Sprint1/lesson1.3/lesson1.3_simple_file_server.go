package main

import (
	"net/http"
)

func main() {
	// простейший сервер, которому доступны все файлы в поддиректории static
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
