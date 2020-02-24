package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
)

func findBiggerSquare(num *big.Int) *big.Int {
	res := new(big.Int)
	res.Add(res.Sqrt(num), big.NewInt(1))
	if res.Cmp(new(big.Int).Sub(num, big.NewInt(1))) != -1 {
		return nil
	}
	return res
}

func simpleSearch(num *big.Int) *big.Int {
	var (
		sqr  = findBiggerSquare(num)
		tmp  = new(big.Int)
		curr = new(big.Int).Sub(tmp.Mul(sqr, sqr), num)
	)

	for {
		s := new(big.Int).Sqrt(curr)
		tmp.Mul(s, s)
		fmt.Printf("sqr: %s\n", sqr)
		sign := tmp.Sub(tmp, curr).Sign()

		if sign == 0 {
			return s
		}
		sqr.Add(sqr, big.NewInt(1))
		curr.Sub(tmp.Mul(sqr, sqr), num)
	}
}

func main() {
	num1Bytes, err := ioutil.ReadFile("test1")
	if err != nil {
		log.Fatal(err)
	}
	num1, _ := new(big.Int).SetString(string(num1Bytes[:len(num1Bytes)-1]), 10)
	res := simpleSearch(num1)
	log.Print(res)
}
