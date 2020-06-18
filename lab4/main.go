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
func correlate(a, b []byte, limit, offset int) float64 {
	min := len(a)
	if len(b) < min {
		min = len(b)
	}

	if offset < 0 {
		offset = 0
	}
	if offset > min {
		fmt.Println("Смещение больше чем размер файла")
		os.Exit(1)
	}

	if limit < min-offset && limit > 0 {
		min = limit + offset
	}

	if offset > min {
		fmt.Println("Смещение больше чем размер файла")
		os.Exit(1)
	}

	// count of matching symbols
	count := 0

	for idx := offset; idx < min; idx++ {
		if a[idx] == b[idx] {
			count++
		}
	}

	return float64(count) / float64(min)
}

func main() {
	limit := flag.Int("n", -1, "количество обрабатываемых символов. -1 ограничивает размер минимальным из текстов")
	offset := flag.Int("s", 0, "смещение относительно начала файла")
	flag.Parse()

	if flag.NArg() != 2 {
		fmt.Println("Необходимо 2 аргумента")
		os.Exit(2)
	}

	a := load(flag.Arg(0))
	b := load(flag.Arg(1))

	num := correlate(a, b, *limit, *offset)

	fmt.Printf("Файлы %q, %q. Ограничение размера: %d, смещение: %d\n",
		flag.Arg(0), flag.Arg(1),
		*limit, *offset)
	fmt.Printf("Отношение совпадающих символов к общему количеству символов: %f\n\n", num)
}
