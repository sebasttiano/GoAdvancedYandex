package main

import (
	"testing"
)

type Operation int

const (
	Add Operation = iota
	Sub
	Mul
	Div
)

type Calculator interface {
	Calculate() int
}

// допишите код
// ...

type Oper struct {
	Type  Operation
	Left  Calculator
	Right Calculator
}

func (o *Oper) Calculate() int {
	switch o.Type {
	case Add:
		return o.Left.Calculate() + o.Right.Calculate()
	case Sub:
		return o.Left.Calculate() - o.Right.Calculate()
	case Mul:
		return o.Left.Calculate() * o.Right.Calculate()
	case Div:
		return o.Left.Calculate() / o.Right.Calculate()
	}
	return 0
}

type Number struct {
	Value int
}

func (n *Number) Calculate() int {
	return n.Value
}

func TestCalc(t *testing.T) {
	root := &Oper{
		Type: Div,
		Left: &Oper{
			Type: Mul,
			Left: &Oper{
				Type:  Add,
				Left:  &Number{Value: 2},
				Right: &Number{Value: 3},
			},
			Right: &Oper{
				Type:  Sub,
				Left:  &Number{Value: 77},
				Right: &Number{Value: 55},
			},
		},
		Right: &Number{Value: 2},
	}
	if root.Calculate() != 55 {
		t.Errorf(`get %d want %d`, root.Calculate(), 77)
	}
}
