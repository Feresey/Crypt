package main

import (
	"encoding/hex"
	"fmt"
	"hash"
	"os"

	"golang.org/x/crypto/sha3"
)

func getHash(h hash.Hash, name string, data []byte) {
	h.Write(data)

	fmt.Printf("%+22s : %s\n", name, hex.EncodeToString(h.Sum(nil)))
}

func main() {
	var data []byte

	if len(os.Args) == 2 {
		data = []byte(os.Args[1])
	}

	getHash(New256(), "keccak 256", data)
	getHash(sha3.NewLegacyKeccak256(), "legacy stl keccak 256", data)
	getHash(sha3.New256(), "stl keccak 256", data)
}
