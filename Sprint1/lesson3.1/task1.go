package main

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"log"
)

type Config2 struct {
	User string `env:"USER"`
}

func main() {
	var cfg Config2
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Current user is %s\n", cfg.User)
}
