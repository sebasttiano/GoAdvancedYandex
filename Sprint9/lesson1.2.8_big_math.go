package main

import (
	"fmt"
	"math/big"
)

func main() {
	dividend := new(big.Int)
	// устанавливаем значение делимого из строки,
	// потому что константа подходящего числового типа невозможна
	// второй аргумент, 10, система счисления
	dividend.SetString("100000000000000000000500", 10)

	divisor := big.NewInt(10)

	quotient := new(big.Int)
	// операция деления
	quotient.Div(dividend, divisor)
	fmt.Println(quotient)
}
