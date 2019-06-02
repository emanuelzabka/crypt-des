#!/bin/bash

echo "Testing -newkey param"

key=$(./crypt-des -newkey)

if [ "$key" == "" ]; then
	(>&2 echo "[error] No key generated.")
	exit 1
fi

if [ ! ${#key} -eq 16 ]; then
	(>&2 echo "[error] Invalid key generated: '$key'.")
	exit 1
fi
