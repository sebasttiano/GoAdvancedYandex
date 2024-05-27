package main

import (
	"bufio"
	"fmt"
	"os"
)

// PrimeToFile записывает в файл fname простые числа,
// которые меньше или равны n.
func PrimeToFile(n int, fname string) error {
	f, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	s := make([]byte, n+1)

	fmt.Fprintf(w, "%d ", 2)
	count := 1 // уже записали число 2
	// будем перебирать только нечётные числа
	for i := 3; i <= n; i += 2 {
		if s[i] == 0 {
			fmt.Fprintf(w, "%d ", i)
			for k := 2 * i; k <= n; k += i {
				s[k] = 1
			}
			count++
			if count == 10 {
				fmt.Fprintf(w, "\r\n")
				count = 0
			}
		}
	}
	w.Flush()

	return err
}
