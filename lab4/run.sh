#!/bin/bash

set -e

function run() {
    "$1" -n 200
    "$1" -n 200  -s 200
    "$1" -n 500
    "$1" -n 1500 -s 1500
    "$1" -n 10000
    "$1" -s 100500
    "$1"
}

function human() {
    ./lab4 "$@" testdata/humanity/{harry,jule}.txt
}
printf "\ncompare human texts\n\n"
run human

function human_symbols() {
    ./lab4 "$@" testdata/humanity/harry.txt testdata/random/symbols_1.txt
}
printf "\ncompare human text & random symbols\n\n"
run human_symbols

function human_words() {
    ./lab4 "$@" testdata/random/words_1.txt testdata/humanity/jule.txt
}
printf "\ncompare human text & random words\n\n"
run human_words

function random_words() {
    ./lab4 "$@" testdata/random/words_{1,2}.txt
}
printf "\ncompare random words\n\n"
run random_words

function random_symbols() {
    ./lab4 "$@" testdata/random/symbols_{1,2}.txt
}
printf "\ncompare random symbols\n\n"
run random_symbols
