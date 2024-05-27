package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

// PrimeToFile1 записывает в файл fname простые числа,
// которые меньше или равны n.
func PrimeToFile1(n int, fname string) error {
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

// SearchPrime находит строки в файле fname с указанными в primes простыми числами.
func SearchPrime(fname string, primes []int) (lines []int, err error) {
	var f *os.File
	f, err = os.Open(fname)
	if err != nil {
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	// можно каждое считанное число переводить в int, но быстрее будет
	// один раз перевести числа в строки и сравнивать именно строки
	sprime := make([]string, len(primes))
	for i, v := range primes {
		sprime[i] = strconv.Itoa(v)
	}
	// результирующий массив с номерами строк
	lines = make([]int, len(primes))
	var (
		iline int // текущая строка
		inum  int // позиция текущего проверяемого числа
	)
	for scanner.Scan() {
		iline++
		nums := strings.Split(strings.TrimRight(scanner.Text(), " "), " ")
		if len(nums) > 0 {
			// для ускорения можно сравнить последнее число в строке
			last := nums[len(nums)-1]
			if len(last) < len(sprime[inum]) ||
				(len(last) == len(sprime[inum]) && last < sprime[inum]) {
				continue
			}
			for _, v := range nums {
				if sprime[inum] == v {
					lines[inum] = iline
					inum++
					if inum == len(lines) {
						return
					}
				}
			}
		}
	}
	err = scanner.Err()
	return
}

func TestPrime1(t *testing.T) {
	if err := PrimeToFile(50000, "prime.txt"); err != nil {
		t.Error(err)
	}
	lines, err := SearchPrime("prime.txt", []int{2, 137, 977, 2239, 2293, 9941,
		16693, 16699, 26647, 37579, 48337})
	if err != nil {
		t.Error(err)
	}
	if fmt.Sprint(lines) != "[1 4 17 34 34 123 193 194 293 398 498]" {
		t.Errorf("unexpected lines %v", lines)
	}
}
