#!/bin/bash

set -e

go build

./lab4 -n 200         testdata/{harry,jule}.txt
./lab4 -n 200  -s 200 testdata/{harry,jule}.txt
./lab4 -n 500         testdata/{harry,jule}.txt
./lab4 -n 1500 -s 500 testdata/{harry,jule}.txt
./lab4 -n 10000       testdata/{harry,jule}.txt
./lab4 -s 100500      testdata/{harry,jule}.txt
