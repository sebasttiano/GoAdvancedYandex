package main

import "fmt"

// интерфейсы
type publisher interface {
	register(observer)
	deregister(observer)
	notify()
}

type observer interface {
	update(string)
	getID() string
}

// реализация publisher
type event struct {
	observers   map[string]observer
	description string
}

func (e *event) register(o observer) {
	if e.observers == nil {
		e.observers = make(map[string]observer)
	}
	e.observers[o.getID()] = o
}

func (e *event) deregister(o observer) {
	delete(e.observers, o.getID())
}

func (e *event) notify() {
	for _, observer := range e.observers {
		observer.update(e.description)
	}
}

func (e *event) update(desc string) {
	e.description = desc
	e.notify()
}

// реализация observer
type subscriber struct {
	id string
}

func (s *subscriber) update(desc string) {
	fmt.Printf("The %s subscriber is notified of the %s event\n", s.id, desc)
}

func (s *subscriber) getID() string {
	return s.id
}

func main() {
	pub := new(event)
	sub1 := subscriber{"first"}
	sub2 := subscriber{"second"}
	sub3 := subscriber{"third"}
	pub.register(&sub1)
	pub.register(&sub2)
	pub.register(&sub3)
	pub.update("Alert")
}
