package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
)

//func PrimeToFile(n int, fname string) error {
//	// функция из задания 1
//}

// MaxPrimeInterval читает файл fname с простыми числами и возвращает
// максимальный интервал dif между числами.
// prime — первое число с найденным интервалом до следующего числа.
func MaxPrimeInterval(fname string) (prime int, dif int, err error) {
	var f *os.File
	f, err = os.Open(fname)
	if err != nil {
		return
	}
	defer f.Close()
	r := bufio.NewReader(f)

	var wnum string
	var ptr []byte
	var prev, num int
	for err != io.EOF {
		ptr, err = r.ReadSlice(' ')
		if err != nil && err != io.EOF {
			return
		}
		wnum = strings.TrimSpace(string(ptr))
		if len(wnum) == 0 {
			continue
		}
		if num, err = strconv.Atoi(wnum); err != nil {
			return
		}
		if num-prev > dif {
			dif = num - prev
			prime = prev
		}
		prev = num
	}
	return prime, dif, nil
}

func TestPrime(t *testing.T) {
	if err := PrimeToFile(50000, "prime.txt"); err != nil {
		t.Error(err)
	}
	if prime, dif, err := MaxPrimeInterval("prime.txt"); err != nil {
		t.Error(err)
	} else if prime != 31397 || dif != 72 {
		t.Errorf("unexpected prime=%d / interval=%d", prime, dif)
	}
}
