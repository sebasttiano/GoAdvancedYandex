package main

import (
	"fmt"
)

// CarPart — семейство типов, которым хотим добавить
// функциональность детали автомобиля.
type CarPart interface {
	Accept(CarPartVisitor)
}

// CarPartVisitor — интерфейс visitor,
// в его коде и содержится новая функциональность.
type CarPartVisitor interface {
	testWheel(wheel *Wheel)
	testEngine(engine *Engine)
}

// Wheel — реализация деталей.
type Wheel struct {
	Name string
}

// Accept — единственный метод, который нужно добавить типам семейства,
// ссылка на метод visitor.
func (w *Wheel) Accept(visitor CarPartVisitor) {
	visitor.testWheel(w)
}

type Engine struct{}

func (e *Engine) Accept(visitor CarPartVisitor) {
	visitor.testEngine(e)
}

type Car2 struct {
	parts []CarPart
}

// NewCar2 — конструктор автомобиля.
func NewCar2() *Car2 {
	this := new(Car2)
	this.parts = []CarPart{
		&Wheel{"front left"},
		&Wheel{"front right"},
		&Wheel{"rear right"},
		&Wheel{"rear left"},
		&Engine{}}
	return this
}

func (c *Car2) Accept(visitor CarPartVisitor) {
	for _, part := range c.parts {
		part.Accept(visitor)
	}
}

// TestVisitor — конкретная реализация visitor,
// которая может проверять колёса и двигатель.
type TestVisitor struct {
}

func (v *TestVisitor) testWheel(wheel *Wheel) {
	fmt.Printf("Testing the %v wheel\n", wheel.Name)
}

func (v *TestVisitor) testEngine(engine *Engine) {
	fmt.Println("Testing engine")
}

func main() {
	// клиентский код
	car := NewCar2()
	visitor := new(TestVisitor)
	car.Accept(visitor)
}

//import "go/ast"
//
//func (v *visitor) Visit(n ast.Node) ast.Visitor {
//	switch x := v.(type)
//	case *ast.Ident:
//		// do stuff with *ast.Ident
//	case *ast.IfStmt:
//		// do stuff with *ast.IfStmt
//		// ...
//	}
