package main

import "fmt"

// Processor — интерфейс обработчика.
type Processor interface {
	Process(Request)
	SetNext(Processor)
}

type Kind int

const (
	Urgent Kind = 1 << iota
	Special
	Valuable
)

// Request описывает поля запроса.
type Request struct {
	Kind Kind
	Data string
}

// Printer2 — обработчик.
type Printer2 struct {
	next Processor
}

func (p *Printer2) Process(r Request) {
	fmt.Printf("Printer2: %s\n", r.Data)
	if p.next != nil {
		p.next.Process(r)
	}
}

func (p *Printer2) SetNext(next Processor) {
	p.next = next
}

// Saver — обработчик.
type Saver struct {
	next Processor
}

func (s *Saver) Process(r Request) {
	// обрабатывает не все запросы
	if r.Kind&(Valuable|Special) != 0 {
		fmt.Printf("Saver: %s\n", r.Data)
		// сохраняем состояние
		// ...
	}
	if s.next != nil {
		s.next.Process(r)
	}
}

func (s *Saver) SetNext(next Processor) {
	s.next = next
}

// Logger — обработчик.
type Logger struct {
	next Processor
}

func (l *Logger) Process(r Request) {
	if r.Kind&Urgent != 0 {
		fmt.Printf("Logger: %s\n", r.Data)
		// записываем в лог
		// ...
	}
	if l.next != nil {
		l.next.Process(r)
	}
}

func (l *Logger) SetNext(next Processor) {
	l.next = next
}

// клиентский код
func main() {
	p := new(Printer2)
	l := new(Logger)
	l.SetNext(p)
	s := new(Saver)
	s.SetNext(l)
	s.Process(Request{0, "Average"})
	s.Process(Request{Valuable, "Do not forget"})
	s.Process(Request{Urgent | Special, "Alert!!!"})
}
