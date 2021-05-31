# Base58 batch key generator

Dummy go program for generating a batch of random keys in base58 format

```
./b58keygen -h
Usage of ./b58keygen:
  -len int
        key length (default 16)
  -size int
        number of keys to generate (default 10)
  -wordList string
        word list to use instead of random bytes

./b58keygen -len=12 -size=2500 > keys.txt
./b58keygen -len=6 -size=500 -wordList=wordlists/catala.txt > mnemonics.cat.csv
```
