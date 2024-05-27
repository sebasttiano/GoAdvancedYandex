package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`i\d+`)
	list := re.FindAllString(`i j34 i78 k56 a3 b78 i0 dd ij125 i534`, -1)
	fmt.Println(list)

	re2 := regexp.MustCompile(`(\w+)\.(\w+)`)
	fmt.Printf("%q\n",
		re2.FindAllStringSubmatch("Для поиска используйте yandex.ru или ya.ru.", -1))
}
