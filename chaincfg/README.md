chaincfg
========

[![Build Status](http://img.shields.io/travis/eacsuite/eacd.svg)]
(https://travis-ci.org/eacsuite/eacd) [![ISC License]
(http://img.shields.io/badge/license-ISC-blue.svg)](http://copyfree.org)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)]
(http://godoc.org/github.com/eacsuite/eacd/chaincfg)

Package chaincfg defines chain configuration parameters for the three standard
Bitcoin networks and provides the ability for callers to define their own custom
Bitcoin networks.

Although this package was primarily written for eacd, it has intentionally been
designed so it can be used as a standalone package for any projects needing to
use parameters for the standard Bitcoin networks or for projects needing to
define their own network.

## Sample Use

```Go
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/eacsuite/eacutil"
	"github.com/eacsuite/eacd/chaincfg"
)

var testnet = flag.Bool("testnet", false, "operate on the testnet Bitcoin network")

// By default (without -testnet), use mainnet.
var chainParams = &chaincfg.MainNetParams

func main() {
	flag.Parse()

	// Modify active network parameters if operating on testnet.
	if *testnet {
		chainParams = &chaincfg.TestNet4Params
	}

	// later...

	// Create and print new payment address, specific to the active network.
	pubKeyHash := make([]byte, 20)
	addr, err := eacutil.NewAddressPubKeyHash(pubKeyHash, chainParams)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(addr)
}
```

## Installation and Updating

```bash
$ go get -u github.com/eacsuite/eacd/chaincfg
```

## GPG Verification Key

All official release tags are signed by Conformal so users can ensure the code
has not been tampered with and is coming from the btcsuite developers.  To
verify the signature perform the following:

- Download the public key from the Conformal website at
  https://opensource.conformal.com/GIT-GPG-KEY-conformal.txt

- Import the public key into your GPG keyring:
  ```bash
  gpg --import GIT-GPG-KEY-conformal.txt
  ```

- Verify the release tag with the following command where `TAG_NAME` is a
  placeholder for the specific tag:
  ```bash
  git tag -v TAG_NAME
  ```

## License

Package chaincfg is licensed under the [copyfree](http://copyfree.org) ISC
License.
