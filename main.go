package main

import (
	"fmt"
	"go-chain/modules"
)

func main() {
	fmt.Println("Creating the go-chain")
	newBlockchain := modules.NewBlockchain()

	newBlockchain.AddBlock("transaction1")
	newBlockchain.AddBlock("transaction2")

	for i, block := range newBlockchain.Blocks {
		fmt.Printf("BlockId: %x\n", i)
		fmt.Printf("Timestamp: %x\n", block.Ts+int64(i))
		fmt.Printf("Hash of the block: %x\n", block.CurrentBlockHash)
		fmt.Printf("Hash of the previous block: %x\n", block.PreviousBlockHash)
		fmt.Printf("All the transactions: %x\n", block.BlockData)
	}

}
