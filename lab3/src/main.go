package main

import (
	"fmt"
	// "github.com/ebfe/keccak"
	// "golang.org/x/crypto/sha3"
)

func main() {
	h := New256()
	// h := sha3.NewLegacyKeccak256()
	// fmt.Println("vim-go")
	_, _ = h.Write([]byte("Милько Павел Алексеевич"))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Println(string(h.Sum(nil)))
}
