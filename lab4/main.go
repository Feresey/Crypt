package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

// load a file to []byte.
// fail on any error
func load(path string) []byte {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Ошибка при считывании файла.\nФайл: %q, ошибка: %v\n", path, err)
		os.Exit(1)
	}

	return data
}

// correlate matching symbols
func correlate(a, b []byte, n int) float64 {
	min := len(a)
	if len(b) < min {
		min = len(b)
	}

	if n < min {
		min = n
	}

	// count of matching symbols
	count := 0

	for idx := 0; idx < min; idx++ {
		if a[idx] == b[idx] {
			count++
		}
	}

	return float64(count) / float64(min)
}

func main() {
	cut := flag.Int("n", -1, "количество обрабатываемых символов. -1 ограничивает размер минимальным из текстов")
	flag.Parse()

	if flag.NArg() != 2 {
		fmt.Println("Необходимо 2 аргумента")
		os.Exit(2)
	}

	a := load(flag.Arg(1))
	b := load(flag.Arg(2))

	num := correlate(a, b, *cut)

	fmt.Printf("Отношение совпадающих символов к общему количеству символов: %f\n", num)
}
