package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

func main() {
	var (
		randomSymbols = flag.Bool("s", false, "generate random symbols")
		randomWords   = flag.Bool("w", false, "generate randowm words")
		wordsList     = flag.String("i", "", "file with words list")
		size          = flag.Int("n", 200, "number of bytes to generate")
		output        = flag.String("o", "", "output file")
	)
	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	if *randomSymbols == *randomWords {
		fmt.Println("required one of parameters: s, w")
		flag.Usage()
		os.Exit(2)
	}

	out, err := os.Create(*output)
	if err != nil {
		fmt.Printf("Error create file. Filename: %q, error: %v\n", *output, err)
		os.Exit(1)
	}
	defer out.Close()

	if *randomSymbols {
		GenRandomSymbols(*size, out)
		return
	}

	// random words

	if *wordsList == "" {
		fmt.Println("required parameter: i")
		flag.Usage()
		os.Exit(2)
	}

	wordsBytes, err := ioutil.ReadFile(*wordsList)
	if err != nil {
		fmt.Printf("Error read file. Filename: %q, error: %v\n", *wordsList, err)
		os.Exit(1)
	}

	words := bytes.Split(wordsBytes, []byte{'\n'})

	GenRandomWords(words, *size, out)
}

func GenRandomSymbols(size int, wr io.Writer) {
	alphabet := make([]byte, 0, 100)
	for i := 'a'; i < 'z'; i++ {
		alphabet = append(alphabet, byte(i))
	}
	for i := 'A'; i < 'Z'; i++ {
		alphabet = append(alphabet, byte(i))
	}
	for i := '0'; i < '9'; i++ {
		alphabet = append(alphabet, byte(i))
	}

	alphabet = append(alphabet, []byte("+-=!?., \n\t;'\":#%*`~/")...)

	w := bufio.NewWriter(wr)

	for ; size > 0; size-- {
		err := w.WriteByte(alphabet[rand.Int31n(int32(len(alphabet)))])
		if err != nil {
			fmt.Printf("Unable write to output file: %v\n", err)
			os.Exit(1)
		}
	}

	if err := w.Flush(); err != nil {
		fmt.Printf("Unable write to output file: %v\n", err)
		os.Exit(1)
	}
}

func GenRandomWords(words [][]byte, size int, w io.Writer) {
	for size > 0 {
		word := words[rand.Int31n(int32(len(words)))]
		size -= len(word) + 1

		_, err := w.Write(append(word, ' '))
		if err != nil {
			fmt.Printf("Unable write to output file: %v\n", err)
			os.Exit(1)
		}
	}
}
