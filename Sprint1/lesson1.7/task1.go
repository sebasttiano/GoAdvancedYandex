package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	response, err := http.Get("https://practicum.yandex.ru")
	if err != nil {
		fmt.Println(err)
	}

	body, err := io.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body[:512]))
}
