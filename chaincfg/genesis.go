// Copyright (c) 2014-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package chaincfg

import (
	"time"

	"github.com/eacsuite/eacd/chaincfg/chainhash"
	"github.com/eacsuite/eacd/wire"
)

// genesisCoinbaseTx is the coinbase transaction for the genesis blocks for
// the main network, regression test network, and test network (version 3).
var genesisCoinbaseTx = wire.MsgTx{
	Version: 2,
	TxIn: []*wire.TxIn{
		{
			PreviousOutPoint: wire.OutPoint{
				Hash:  chainhash.Hash{},
				Index: 0xffffffff,
			},
			SignatureScript: []byte{
				0x04, 0xff, 0xff, 0x00, 0x1d, 0x01, 0x04, 0x40, 0x4e, 0x59, 0x20, 0x54, 0x69, 0x6d, 0x65, 0x73, // |.......@NY Times|
				0x20, 0x30, 0x35, 0x2f, 0x4f, 0x63, 0x74, 0x2f, 0x32, 0x30, 0x31, 0x31, 0x20, 0x53, 0x74, 0x65, // | 05/Oct/2011 Ste|
				0x76, 0x65, 0x20, 0x4a, 0x6f, 0x62, 0x73, 0x2c, 0x20, 0x41, 0x70, 0x70, 0x6c, 0x65, 0xe2, 0x80, // |ve Jobs, Apple..|
				0x99, 0x73, 0x20, 0x56, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x2c, 0x20, 0x44, 0x69, // |.s Visionary, Di|
				0x65, 0x73, 0x20, 0x61, 0x74, 0x20, 0x35, 0x36, // |es at 56|

			},
			Sequence: 0xffffffff,
		},
	},
	TxOut: []*wire.TxOut{
		{
			Value: 0x12a05f200,
			PkScript: []byte{
				0x41, 0x4, 0x1, 0x84, 0x71, 0xf, 0xa6, 0x89,
				0xad, 0x50, 0x23, 0x69, 0xc, 0x80, 0xf3, 0xa4,
				0x9c, 0x8f, 0x13, 0xf8, 0xd4, 0x5b, 0x8c, 0x85,
				0x7f, 0xbc, 0xbc, 0x8b, 0xc4, 0xa8, 0xe4, 0xd3,
				0xeb, 0x4b, 0x10, 0xf4, 0xd4, 0x60, 0x4f, 0xa0,
				0x8d, 0xce, 0x60, 0x1a, 0xaf, 0xf, 0x47, 0x2,
				0x16, 0xfe, 0x1b, 0x51, 0x85, 0xb, 0x4a, 0xcf,
				0x21, 0xb1, 0x79, 0xc4, 0x50, 0x70, 0xac, 0x7b,
				0x3, 0xa9, 0xac,
			},
		},
	},
	LockTime: 0,
}

// genesisHash is the hash of the first block in the block chain for the main
// network (genesis block).
var genesisHash = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0xed, 0x6c, 0x14, 0xf9, 0xe2, 0x11, 0x95, 0x6c,
	0xcc, 0xd8, 0xbc, 0x8b, 0x72, 0x06, 0xad, 0x83,
	0xe4, 0x18, 0xf7, 0x9a, 0xcb, 0xf1, 0x38, 0x05,
	0x1c, 0x30, 0x03, 0xf4, 0x4d, 0x7d, 0x71, 0x21, 
 })

// genesisMerkleRoot is the hash of the first transaction in the genesis block
// for the main network.
var genesisMerkleRoot = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0xb1, 0x4b, 0x5c, 0x80, 0x81, 0x6e, 0x64, 0xe6,
	0x16, 0xa2, 0xab, 0xb5, 0xe5, 0xb0, 0x39, 0x63, 
	0x94, 0x81, 0x7f,  0x4d, 0xf0, 0xc1, 0x2a, 0x45, 
	0x91, 0x18, 0x41, 0x10, 0x36, 0x7c, 0x75, 0x13, 
 })



// genesisBlock defines the genesis block of the block chain which serves as the
// public transaction ledger for the main network.
var genesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  chainhash.Hash{},         // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: genesisMerkleRoot,        // 13757c3610411891452ac1f04d7f81946339b0e5b5aba216e6646e81805c4bb1
		Timestamp:  time.Unix(1386746168, 0),
		Bits:       0x1e0ffff0,
		Nonce:      12468024,
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}

// regTestGenesisHash is the hash of the first block in the block chain for the
// regression test network (genesis block).
var regTestGenesisHash = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0xf9, 0x16, 0xc4, 0x56, 0xfc, 0x51, 0xdf, 0x62,
	0x78, 0x85, 0xd7, 0xd6, 0x74, 0xed, 0x02, 0xdc,
	0x88, 0xa2, 0x25, 0xad, 0xb3, 0xf0, 0x2a, 0xd1,
	0x3e, 0xb4, 0x93, 0x8f, 0xf3, 0x27, 0x08, 0x53,
})

// regTestGenesisMerkleRoot is the hash of the first transaction in the genesis
// block for the regression test network.  It is the same as the merkle root for
// the main network.
var regTestGenesisMerkleRoot = genesisMerkleRoot

// regTestGenesisBlock defines the genesis block of the block chain which serves
// as the public transaction ledger for the regression test network.
var regTestGenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  chainhash.Hash{},         // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: regTestGenesisMerkleRoot, // 97ddfbbae6be97fd6cdf3e7ca13232a3afff2353e29badfab7f73011edd4ced9
		Timestamp:  time.Unix(1386746169, 0), // 2011-02-02 23:16:42 +0000 UTC
		Bits:       0x1e0ffff0,               // 545259519 [7fffff0000000000000000000000000000000000000000000000000000000000]
		Nonce:      12468025,
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}

// testNet4GenesisHash is the hash of the first block in the block chain for the
// test network (version 4).
var testNet4GenesisHash = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0xa0, 0x29, 0x3e, 0x4e, 0xeb, 0x3d, 0xa6, 0xe6,
	0xf5, 0x6f, 0x81, 0xed, 0x59, 0x5f, 0x57, 0x88,
	0xd, 0x1a, 0x21, 0x56, 0x9e, 0x13, 0xee, 0xfd,
	0xd9, 0x51, 0x28, 0x4b, 0x5a, 0x62, 0x66, 0x49,
})

// testNet3GenesisMerkleRoot is the hash of the first transaction in the genesis
// block for the test network (version 3).  It is the same as the merkle root
// for the main network.
var testNet3GenesisMerkleRoot = genesisMerkleRoot

// testNet4GenesisMerkleRoot is the hash of the first transaction in the genesis
// block for the test network (version 4).  It is the same as the merkle root
// for the main network.
var testNet4GenesisMerkleRoot = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0xb1, 0x4b, 0x5c, 0x80, 0x81, 0x6e, 0x64, 0xe6,
	0x16, 0xa2, 0xab, 0xb5, 0xe5, 0xb0, 0x39, 0x63, 
	0x94, 0x81, 0x7f,  0x4d, 0xf0, 0xc1, 0x2a, 0x45, 
	0x91, 0x18, 0x41, 0x10, 0x36, 0x7c, 0x75, 0x13, 
})

// testNet4GenesisBlock defines the genesis block of the block chain which
// serves as the public transaction ledger for the test network (version 4).
var testNet4GenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:   2,
		PrevBlock:  chainhash.Hash{},          // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: testNet4GenesisMerkleRoot, // 97ddfbbae6be97fd6cdf3e7ca13232a3afff2353e29badfab7f73011edd4ced9
		Timestamp:  time.Unix(1386746169, 0),
		Bits:       0x1e0ffff0,
		Nonce:      12468025,
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}

// simNetGenesisHash is the hash of the first block in the block chain for the
// simulation test network.
var simNetGenesisHash = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0xbe, 0xa4, 0x3b, 0x3a, 0xb4, 0xbc, 0x9d, 0x86,
	0x2b, 0xfa, 0xbc, 0x0d, 0x45, 0xa3, 0xc5, 0xc9,
	0xd0, 0x9f, 0x66, 0x20, 0xce, 0x7d, 0xb6, 0x78,
	0xb5, 0xdc, 0x1d, 0xb8, 0x19, 0xc9, 0x35, 0x12,
})

// simNetGenesisMerkleRoot is the hash of the first transaction in the genesis
// block for the simulation test network.  It is the same as the merkle root for
// the main network.
var simNetGenesisMerkleRoot = genesisMerkleRoot

// simNetGenesisBlock defines the genesis block of the block chain which serves
// as the public transaction ledger for the simulation test network.
var simNetGenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  chainhash.Hash{},         // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: simNetGenesisMerkleRoot,  // 4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b
		Timestamp:  time.Unix(1401292357, 0), // 2014-05-28 15:52:37 +0000 UTC
		Bits:       0x207fffff,               // 545259519 [7fffff0000000000000000000000000000000000000000000000000000000000]
		Nonce:      2,
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}
