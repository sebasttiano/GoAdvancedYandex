package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"strings"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func main() {
	var users []User
	url := "https://jsonplaceholder.typicode.com/users"

	client := resty.New()

	_, err := client.R().
		SetResult(&users).
		Get(url)
	if err != nil {
		log.Printf("Error occured: %v", err)
	}

	var out []string
	for _, v := range users {
		out = append(out, v.Username)
	}
	fmt.Println(strings.Join(out, ` `))
	// если выбрали resty, используйте SetResult(&users)
	// для получения результата сразу в виде массива
	// ...
}
