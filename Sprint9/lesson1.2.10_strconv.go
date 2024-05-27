package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Вводите числа, они будут суммироваться:")
	summ := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Ошибка ввода. Попробуйте ещё раз:")
		}
		summ += number
		fmt.Printf("Сумма составила %v\nВведите число:", summ)
	}
}
