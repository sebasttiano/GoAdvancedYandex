package main

import (
	"errors"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

// --addr=example.com:60
type NetAddress struct {
	Host string
	Port int
}

func (n *NetAddress) String() string {
	return fmt.Sprintf("%s:%d", n.Host, n.Port)
}

func (n *NetAddress) Set(flagValue string) error {
	if val := strings.Split(flagValue, ":"); len(val) == 2 {
		n.Host = val[0]
		n.Port, _ = strconv.Atoi(val[1])
	} else {
		errors.New("address must be in a format host:port")
	}
	return nil
}

func main() {
	addr := new(NetAddress)
	// если интерфейс не реализован,
	// здесь будет ошибка компиляции
	_ = flag.Value(addr)
	// проверка реализации
	flag.Var(addr, "addr", "Net address host:port")
	flag.Parse()
	fmt.Println(addr.Host)
	fmt.Println(addr.Port)
}
