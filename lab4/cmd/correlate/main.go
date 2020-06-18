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
func correlate(a, b []byte, limit, offset int) (count int, size int) {
	size = len(a)
	if len(b) < size {
		size = len(b)
	}

	if offset < 0 {
		offset = 0
	}
	if offset > size {
		fmt.Println("Смещение больше чем размер файла")
		os.Exit(1)
	}

	if limit < size-offset && limit > 0 {
		size = limit + offset
	}

	if offset > size {
		fmt.Println("Смещение больше чем размер файла")
		os.Exit(1)
	}

	for idx := offset; idx < size; idx++ {
		if a[idx] == b[idx] {
			count++
		}
	}

	return count, size
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

	count, size := correlate(a, b, *limit, *offset)

	fmt.Printf(`%d & %d & %f\\ \hline`+"\n", size-*offset, *offset, float64(count)/float64(size))
}

// 	fmt.Printf("Файлы %q, %q. Ограничение размера: %d, смещение: %d\n",
// 		flag.Arg(0), flag.Arg(1),
// 		size, *offset)
// 	fmt.Printf("Отношение совпадающих символов к общему количеству символов: %f\n\n",
// 		float64(count)/float64(size))
// }
