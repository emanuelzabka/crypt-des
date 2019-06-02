#!/bin/bash

echo "Testing encryption/decryption through standard input/output"

key=$(./crypt-des -newkey)
input="ABCDEFGHIJKLM012345678"
output=$(echo -n $input | ./crypt-des -encrypt -k $key | ./crypt-des -decrypt -k $key)
if [ "$input" != "$output" ]; then
	(>&2 echo "[error] Invalid decryption")
	exit 1
fi
