#!/bin/bash

go build -o generate ./cmd/generate

./generate -s -n 620000 -o ./testdata/random/symbols_1.txt
./generate -s -n 620000 -o ./testdata/random/symbols_2.txt

./generate -w -i testdata/random/dict.txt -n 620000 -o ./testdata/random/words_1.txt
./generate -w -i testdata/random/dict.txt -n 620000 -o ./testdata/random/words_2.txt

go build -o lab4 ./cmd/correlate