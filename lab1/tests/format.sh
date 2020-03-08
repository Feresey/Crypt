#!/usr/bin/env bash

[[ $# != 1 ]] && printf "usage:\n./format.sh <file>\n" && exit 1

sed -re 's/([0-9]{80})/\1\n/g' "$1" > "$1".format
