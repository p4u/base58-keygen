package main

import (
	"bytes"
	"crypto/rand"
	"flag"
	"io/ioutil"

	"github.com/btcsuite/btcutil/base58"
)

func main() {
	size := flag.Int("size", 10, "number of keys to generate")
	output := flag.String("output", "keys.csv", "output file")
	len := flag.Int("len", 16, "key length")
	flag.Parse()

	keys := bytes.Buffer{}
	for i := 0; i < *size; i++ {
		keys.WriteString(base58.Encode(RandomBytes(*len)))
		keys.WriteString("\n")
	}
	if err := ioutil.WriteFile(*output, keys.Bytes(), 0o600); err != nil {
		panic(err)
	}
}

func RandomBytes(n int) []byte {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return b
}
