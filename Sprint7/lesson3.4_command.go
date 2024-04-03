package main

import "fmt"

type receiver struct {
	args string
}

func (r *receiver) execute() {
	fmt.Println(r.args)
}

var concrete = receiver{"John Doe"}

var command = concrete.execute

func main() {
	command()
}
