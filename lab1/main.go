package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math/big"
	"os"
)

var (
	one = big.NewInt(1)
	two = big.NewInt(2)
)

func findBiggerSquare(num *big.Int) *big.Int {
	res := new(big.Int)
	res.Add(res.Sqrt(num), one)
	if res.Cmp(new(big.Int).Sub(num, one)) != -1 {
		return nil
	}
	return res
}

const (
	less = iota - 1
	equal
	more
)

func isSqaure(num *big.Int) bool {
	tmp := new(big.Int)
	tmp.Sqrt(num)
	tmp.Mul(tmp, tmp)

	return tmp.Cmp(num) == equal
}

func simpleSearch(num *big.Int) (*big.Int, *big.Int) {
	k := findBiggerSquare(num)
	if k == nil {
		return num, one
	}

	var (
		prev = new(big.Int).Sub(new(big.Int).Mul(k, k), num)
		max  = new(big.Int).Mul(num, num)
	)

	for prev.Cmp(max) != more {
		if isSqaure(prev) {
			y := new(big.Int)
			y.Sqrt(num)

			x := new(big.Int).Sub(k, y)

			return x, y.Add(k, y)
		}

		prev.Add(prev, k)
		prev.Add(prev, k)
		prev.Add(prev, one)

		k.Add(k, one)
	}

	return num, one
}

func main() {
	var (
		input string
		in    io.Reader
		data  string
	)
	flag.StringVar(&input, "i", "-", "Файл с входными данными")
	flag.Parse()

	if input == "-" {
		in = os.Stdin
	} else {
		var err error
		in, err = os.Open(input)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Fscan(in, &data)

	fmt.Printf("%q\n", data)
	num1, _ := new(big.Int).SetString(data, 10)
	fmt.Println(num1)
	a, b := simpleSearch(num1)
	fmt.Printf("%s\n%s\n", a, b)
}
