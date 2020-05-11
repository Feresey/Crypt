package main

import (
	"encoding/hex"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/sha3"
)

func TestMyHash(t *testing.T) {
	tests := [][]byte{
		[]byte(""),
		{0b11001},
		[]byte("abc"),
		[]byte("The quick brown fox jumps over the lazy dog"),
		[]byte("LAIHdibisueo462894381b19 g g1re1798236t \u004231"),
	}

	// test1600 := make([]byte, 200)

	// for i := 0; i < 200; i++ {
	// 	test1600[i] = 0b11000101
	// }

	// tests = append(tests, test1600)

	for _, tt := range tests {
		my := New256()
		stl := sha3.New256()

		_, err := my.Write(tt)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		_, err = stl.Write(tt)
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		got := strings.ToUpper(hex.EncodeToString(my.Sum(nil)))
		want := strings.ToUpper(hex.EncodeToString(stl.Sum(nil)))

		assert.Equal(t,
			want,
			got,
		)
		fmt.Printf("hash: %q, text %q\n", got, tt)
	}
}
