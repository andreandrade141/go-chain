package modules

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Ts                int64
	PreviousBlockHash []byte
	CurrentBlockHash  []byte
	BlockData         []byte
}

func (block *Block) SetHash() {
	ts := []byte(strconv.FormatInt(block.Ts, 10))
	headers := bytes.Join([][]byte{ts, block.PreviousBlockHash, block.BlockData}, []byte{})
	hash := sha256.Sum256(headers)
	block.CurrentBlockHash = hash[:]
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), prevBlockHash, []byte{}, []byte(data)}
	block.SetHash()
	return block
}

func GenesisBlock() *Block {
	return NewBlock("genesis", []byte{})
}
