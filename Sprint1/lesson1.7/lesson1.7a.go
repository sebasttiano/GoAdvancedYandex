package main

import (
	"fmt"
	"net/http"
)

func main() {

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println(req.URL)
			return nil
		},
	}
	response, err := client.Get("https://google.com")
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

}
