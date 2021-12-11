package main

import (
	"bufio"
	"bytes"
	"crypto/rand"
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/btcsuite/btcutil/base58"
)

func main() {
	size := flag.Int("size", 10, "number of keys to generate")
	klen := flag.Int("len", 16, "key length")
	wl := flag.String("wordList", "", "word list to use instead of random bytes")
	flag.Parse()

	keys := bytes.Buffer{}

	if *wl != "" {
		file, err := os.Open(*wl)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		list := []string{}
		for scanner.Scan() {
			list = append(list, scanner.Text())
		}

		for i := 0; i < *size; i++ {
			for j := 0; j < *klen; j++ {
				keys.WriteString(list[RandomInt(0, len(list)-1)])
				if j != *klen-1 {
					keys.WriteString(",")
				}
			}
			if i < (*size - 1) {
				keys.WriteString("\n")
			}
		}
	} else {

		for i := 0; i < *size; i++ {
			keys.WriteString(base58.Encode(RandomBytes(*klen))[:*klen])
			if i < (*size - 1) {
				keys.WriteString("\n")
			}
		}
	}
	fmt.Println(keys.String())
}

func RandomBytes(n int) []byte {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return b
}

func RandomInt(min, max int) int {
	num, err := rand.Int(rand.Reader, big.NewInt(int64(max-min)))
	if err != nil {
		panic(err)
	}
	return int(num.Int64()) + min
}
