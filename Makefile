.PHONY: clean build test
SHELL = /bin/bash

build: crypt-des

crypt-des: main.go des/des.go
	go build -o crypt-des main.go

test: crypt-des
	tests/newkey.sh
	tests/newkey_3des.sh
	tests/std_input_output.sh
	tests/std_input_output_3des.sh
	tests/random_encrypt.sh
	tests/random_encrypt_3des.sh

clean:
	rm -f tests/tmp/*
	rm -f crypt-des
