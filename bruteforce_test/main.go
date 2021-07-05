package main

import (
	"bufio"
	"flag"
	"io/ioutil"
	"os"
	"time"

	"github.com/btcsuite/btcutil/base58"
	"go.vocdoni.io/dvote/censustree/gravitontree"
	"go.vocdoni.io/dvote/crypto/ethereum"
	"go.vocdoni.io/dvote/log"
	"go.vocdoni.io/dvote/util"
)

func main() {
	cfile := flag.String("census", "", "census file")
	size := flag.Int("size", 10, "key size")
	flag.Parse()
	log.Init("debug", "stdout")
	stdir, err := ioutil.TempDir("", "")
	if err != nil {
		log.Fatal(err)
	}
	treeName := util.RandomHex(12)
	tree, err := gravitontree.NewTree(treeName, stdir)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(*cfile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		if err := tree.Add(ethereum.HashRaw(scanner.Bytes()), nil); err != nil {
			log.Fatal(err)
		}
	}

	log.Infof("starting bruteforce attack")
	var i int64
	timer := time.Now()
	for {

		mp, err := tree.GenProof(
			ethereum.HashRaw([]byte(base58.Encode(util.RandomBytes(*size)))),
			nil)
		if err != nil {
			log.Fatal(err)
		}
		if mp != nil {
			log.Infof("colision found!!! on %d", i)
		}
		if i%10000 == 0 {
			time.Since(timer).Seconds()
			log.Infof("key %d at %.2f keys/second", i, 10000/time.Since(timer).Seconds())
			timer = time.Now()
		}
		i++
	}
}
