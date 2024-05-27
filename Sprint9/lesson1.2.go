package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

func main() {
	users := []User{
		{"Иванов Иван", 34},
		{"Кузнецов Алексей", 28},
		{"Иванов Иван", 18},
		{"Пешкова Мария", 25},
	}
	fmt.Println(users)
	sort.Slice(users, func(i, j int) bool {
		if users[i].Name == users[j].Name {
			// если имена одинаковые, то сортирует по возрасту
			return users[i].Age < users[j].Age
		}
		return users[i].Name < users[j].Name
	})
	fmt.Print(users)
}
