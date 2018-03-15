package main

type BlockChain struct {
	Blocks []*Block
}

func (bc *BlockChain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.PrevBlockHash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

func NewGenesisBlock()  *Block{
	return NewBlock("Genesis Block", []byte{})
}

func NewBlockChain() *BlockChain{
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}

