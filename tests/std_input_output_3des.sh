#!/bin/bash

echo "Testing -3des encryption/decryption through standard input/output"

key=$(./crypt-des -newkey -3des)
input="ABCDEFGHIJKLM012345678"
output=$(echo -n $input | ./crypt-des -encrypt -3des -k $key | ./crypt-des -decrypt -3des -k $key)
if [ "$input" != "$output" ]; then
	(>&2 echo "[error] Invalid decryption")
	exit 1
fi
