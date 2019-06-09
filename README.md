# crypt-des

Basic implementation of the Data Encryption Standard (DES) cipher algorithm in GOLANG.

Created for educational purposes, not optmized for large inputs.

## Usage

### Creating a key
```
./crypt-des -newkey
```
Note: If you want to use Triple DES, add the param "-3des". Example:
```
./crypt-des -newkey -3des
```
It will output a string with the three keys separated by a colon `:` character.

### Encrypting from standard input to standard output
```
cat originalfile | ./crypt-des -encrypt -k <key> > encryptedfile
```
If the key parameter is not informed, crypt-des generates a random key and outputs it in standard error.
```
cat originalfile | ./crypt-des -encrypt > encryptedfile
Using key: db1e839774ccecdb
```

### Decrypting from standard input to standard output
```
cat encryptedfile | ./crypt-des -decrypt -k <key>
```
### Encrypting from file
```
./crypt-des -encrypt -k <key> -i originalfile -o encryptedfile
```
### Decrypt from file
```
./crypt-des -decrypt -k <key> -i encryptedfile -o originalfile
```
