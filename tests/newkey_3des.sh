#!/bin/bash

echo "Testing -newkey -3des param"

key=$(./crypt-des -newkey -3des)

if [ "$key" == "" ]; then
	(>&2 echo "[error] No key generated.")
	exit 1
fi

if [ ! ${#key} -eq 50 ]; then
	(>&2 echo "[error] Invalid key generated: '$key'.")
	exit 1
fi
