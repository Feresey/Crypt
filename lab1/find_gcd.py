#!/usr/bin/env python

import sys

if len(sys.argv) != 2:
    print("expected filename")
    sys.exit(1)

FILE = open(sys.argv[1], "r")


def read_one(file=FILE):
    try:
        n1 = int(file.readline()[3:])
        n2 = int(file.readline()[3:])
        return (n1, n2)
    except:
        return (0, 0)


def NOD(a, b):
    while a != 0 and b != 0:
        if a > b:
            a %= b
        else:
            b %= a

    return a+b


nums = []

while True:
    n1, n2 = read_one()
    if n1 == 0:
        break
    nums.append((n1, n2))

most = nums[12]
del nums[12]

for num in nums:
    n1 = NOD(most[1], num[0])
    n2 = NOD(most[1], num[1])
    if n1 != 1:
        print(f"NOD:\n{n1}\norig:\n{most[1]}\nsecond:\n{num[0]}\n")
    if n2 != 1:
        print(f"NOD:\n{n2}\norig:\n{most[1]}\nsecond:\n{num[1]}\n")
