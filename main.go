package main

import "fmt"

func main() {
	bc := NewBlockChain()

	bc.AddBlock("Send 1 BTC to baby")
	bc.AddBlock("Send 0.5 BTC to my brother")

	for _, block := range bc.Blocks {
		fmt.Println("=============================================")
		fmt.Printf("Prev.hash : %x\n", block.PrevBlockHash)
		fmt.Printf("Data : %s\n", block.Data)
		fmt.Printf("Hash : %x\n", block.Hash)
	}
}
