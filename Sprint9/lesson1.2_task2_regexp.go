package main

import (
	"fmt"
	"regexp"
)

func main() {
	emails := []string{"$€§@yandex.com", "ivan@mail.ru", "john@gmailyahoo", "fedor@gmail.com", "stepan@yahoo.com", "commanderpike@gmail.com", "greta@abcd@gmail_yahoo.com"}
	re := regexp.MustCompile(`\w+@.+\..+`)
	for _, email := range emails {
		fmt.Println(re.FindAllString(email, -1))
	}
}
