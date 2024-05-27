package main

import (
	"fmt"
	"math"
)

func main() {
	res := math.Sqrt(14.28)
	fmt.Println(res) // 3.778888725538237
	res = math.Sqrt(-10)
	fmt.Println(res) // NaN
	// строка выше выведет NaN, NotaNumber,
	// потому что пытаемся извлечь корень из отрицательного числа,
	// значение NaN можно проверить
	fmt.Println(math.IsNaN(res)) // true
	// в пакете предусмотрены значения +Inf (positive infinity) и -Inf (negative infinity)
	// для краевых значений математических функций
	res = 10 / math.Log(1)
	fmt.Println(res)                // +Inf
	fmt.Println(math.IsInf(res, 1)) // true
}
