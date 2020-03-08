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


def main():
    nums = []
    num_variant = 12

    while True:
        n1, n2 = read_one()
        if n1 == 0:
            break

        nums.append((n1, n2))

    most = nums[num_variant]

    for idx, num in enumerate(nums):
        n1 = NOD(most[1], num[0])
        n2 = NOD(most[1], num[1])
        for id_res, num_nod in enumerate((n1, n2)):
            if num_nod != 1:
                print(
                    f"variant={idx}\nNOD:\n{num_nod}\norig:\n{most[1]}\nsecond:\n{num[id_res]}\n")
                return


if __name__ == "__main__":
    main()
