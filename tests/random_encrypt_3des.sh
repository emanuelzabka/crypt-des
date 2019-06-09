#!/bin/bash

echo "Testing -3des random encryptions"

testEncrypt() {
	filename="tests/tmp/$1"
	echo "Testing encryption of file $filename"
	key=$(./crypt-des -newkey -3des)
	./crypt-des -encrypt -3des -i $filename -o $filename.enc -k $key
	if [ ! -s "$filename.enc" ]; then
		(>&2 echo "[error] Failed to encrypt $filename with key $key")
		exit 1
	fi
	./crypt-des -decrypt -3des -i $filename.enc -o $filename.dec -k $key
	if [ ! -s "$filename.dec" ]; then
		(>&2 echo "[error] Failed to decrypt $filename with key $key")
		exit 1
	fi
	cmp --silent $filename $filename.dec
	if [ $? -ne 0 ]; then
		(>&2 echo "[error] Source file $filename not the same as $filename.dec")
		exit 1
	fi
}

rm tests/tmp/* 2>/dev/null
dd bs=1 count=1 if=/dev/urandom of=tests/tmp/random1 2>/dev/null
testEncrypt random1
dd bs=1 count=7 if=/dev/urandom of=tests/tmp/random2 2>/dev/null
testEncrypt random2
dd bs=1 count=8 if=/dev/urandom of=tests/tmp/random3 2>/dev/null
testEncrypt random3
dd bs=1 count=13 if=/dev/urandom of=tests/tmp/random4 2>/dev/null
testEncrypt random4
dd bs=1 count=20 if=/dev/urandom of=tests/tmp/random5 2>/dev/null
testEncrypt random5
dd bs=32 count=1 if=/dev/urandom of=tests/tmp/random6 2>/dev/null
testEncrypt random6
dd bs=1024 count=1 if=/dev/urandom of=tests/tmp/random7 2>/dev/null
testEncrypt random7
dd bs=1048576 count=2 if=/dev/urandom of=tests/tmp/random8 2>/dev/null
testEncrypt random8
rm tests/tmp/* 2>/dev/null
