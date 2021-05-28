# Base58 batch key generator

Dummy go program for generating a batch of random keys in base58 format

```
./b58keygen -h
Usage of ./b58keygen:
  -len int
        key length (default 16)
  -output string
        output file (default "keys.csv")
  -size int
        number of keys to generate (default 10)


./b58keygen -len=12 -size=2500 -output=keys.csv
```
