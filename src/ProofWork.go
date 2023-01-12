package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math/big"
	"time"
)

type Block struct {
	// meta data
	Timestamp int64
	PrevHash  []byte
	Nonce     int

	//data
	Data []byte

	// hashed Fields used for hashing and mining.
	Hash []byte
}

// Here, targetBits = 16, then target value will have 16 leading zeroes. This means, that we need to find a nonce of
// block hash that as 16 leading zeroes.
const targetBits = 16

func main() {
	// The purpose of this is to find the target value which isnt impossible. The puzzle in PoW, is we find the nonce of
	// block hash that is less than the target value. The difficulty of this task is dependent on how many leading zeroes
	// in the target value. controlled by targetBits.
	var target big.Int

	// here target value is a 256bit integer, and we are shifting left by targetBits.
	target.Lsh(big.NewInt(1), 256-targetBits)

	fmt.Printf("Mining a block...\n")
	start := time.Now()
	block := createBlock(0)
	block.mine(targetBits)

	elapsed := time.Since(start)
	fmt.Printf("Mining took %s\n", elapsed)
}

func createBlock(nonce int) *Block {
	return &Block{
		Timestamp: time.Now().Unix(),
		Data:      []byte("Hello, world!"),
		PrevHash: []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		Nonce: nonce,
		Hash:  []byte{},
	}
}

func (b *Block) hash() *big.Int {
	// compute and return the hash of the block
	hash := sha256.Sum256(b.serialize())

	//This is done because the big.Int type provides a convenient way to represent large integers and perform arithmetic
	//operations on them, while the byte array representation of the hash value may not be as convenient to work with.
	var result big.Int
	result.SetBytes(hash[:])
	return &result
}

func (b *Block) mine(targetBits uint) {
	// Compute the target value
	var target big.Int
	target.Lsh(big.NewInt(1), 256-targetBits)

	// Mine the block by finding a valid nonce
	for {
		hash := b.hash()
		if hash.Cmp(&target) == -1 {
			b.Hash = hash.Bytes()
			break
		}
		b.Nonce++
	}
}

func (b *Block) serialize() []byte {
	var result bytes.Buffer
	if err := binary.Write(&result, binary.BigEndian, b); err != nil {
		// handle error
		panic(err)
	}
	return result.Bytes()
}
