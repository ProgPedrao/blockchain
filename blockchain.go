package main

import (
)

type BlockChain struct{
	blocks []*Block
}

func NewBlockchain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}

func (bc *BlockChain) AddBlock(data string)  {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	NewBlock := NewBlock(data, prevBlock.Hash)

	bc.blocks = append(bc.blocks, NewBlock)
}