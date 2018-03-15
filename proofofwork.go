package main

import (
	"math/big"
	"bytes"
	"fmt"
	"crypto/sha256"
	"math"
)

var (
	//对循环进行限制
	maxNonce = math.MaxInt64
)

//难度  区块头  开头有多少个0
const targetBits = 24

type ProofOfWrok struct {
	block 	*Block
	target 	*big.Int //大整数，将hash转换成大整数与target比较  hash比target小 则挖矿成功
}

func NewProofOfWork(b *Block) *ProofOfWrok{
	//将 big.Int 初始化为 1
	target := big.NewInt(1)

	//左移 256 - targetBits 位。256 是一个 SHA-256 哈希的位数，我们将要使用的是 SHA-256 哈希算法
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofOfWrok{b, target}
	return pow
}

func (pow *ProofOfWrok) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)
	return data
}

func (pow *ProofOfWrok) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	for nonce < maxNonce  {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			fmt.Printf("\r%x", hash)
			break
		} else {
			nonce++
		}
	}
	fmt.Printf("\n\n")
	return nonce, hash[:]
}

func (pow *ProofOfWrok) Validate() bool {
	var hashInt big.Int
	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}

