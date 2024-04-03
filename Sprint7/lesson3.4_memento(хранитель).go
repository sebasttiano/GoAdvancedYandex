package main

import "fmt"

type originator struct {
	state string
}

// createMemento создаёт снимок состояния объекта.
func (e *originator) createMemento() *memento {
	return &memento{state: e.state}
}

// restoreMemento восстанавливает состояние объекта.
func (e *originator) restoreMemento(m *memento) {
	e.state = m.getSavedState()
}

func (e *originator) doSomething(s string) {
	e.state = s
}

func (e *originator) getState() string {
	return e.state
}

type memento struct {
	state string
}

func (m *memento) getSavedState() string {
	return m.state
}

type caretaker struct {
	mementos []*memento
}

// addMemento добавляет снимок состояния.
func (c *caretaker) addMemento(m *memento) {
	c.mementos = append(c.mementos, m)
}

// getMemento возвращает снимок состояния.
func (c *caretaker) getMemento(index int) *memento {
	return c.mementos[index]
}

func main() {

	caretaker := &caretaker{
		mementos: make([]*memento, 0),
	}

	originator := &originator{
		state: "A",
	}

	fmt.Printf("Current state: %s\n", originator.getState())
	caretaker.addMemento(originator.createMemento())

	originator.doSomething("B")
	fmt.Printf("Current state: %s\n", originator.getState())
	caretaker.addMemento(originator.createMemento())

	originator.doSomething("C")
	fmt.Printf("Current state: %s\n", originator.getState())
	caretaker.addMemento(originator.createMemento())

	// восстанавливаем состояния
	originator.restoreMemento(caretaker.getMemento(1))
	fmt.Printf("Restored to: %s\n", originator.getState())

	originator.restoreMemento(caretaker.getMemento(0))
	fmt.Printf("Restored to: %s\n", originator.getState())

}
