SHELL = /bin/bash

build: crypt-des

crypt-des: main.go des/des.go
	go build -o crypt-des main.go

test: crypt-des
	tests/newkey.sh
	tests/std_input_output.sh
	tests/random_encrypt.sh

clean:
	rm -f tests/tmp/*
	rm -f crypt-des
