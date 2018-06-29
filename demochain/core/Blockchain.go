package core

import (
	"log"
	"fmt"
)

type Blockchain struct {
	Blocks []*Block
}

// 创建新区块链
func NewBlockchain() *Blockchain {
	genesisBlock := GenerateGenesisBlock()
	blockchain := Blockchain{}
	blockchain.AppendBlock(&genesisBlock)
	return &blockchain
}

// 向区块链添加区块信息
func (bc *Blockchain) SendData (data string) {
	preBlock := bc.Blocks[len(bc.Blocks) -1]
	newBlock := GenerateNewBlock(*preBlock, data)
	bc.AppendBlock(&newBlock)
}


// 添加新区块
func (bc *Blockchain) AppendBlock(newBlock *Block) {
	if len(bc.Blocks) == 0 {
		bc.Blocks = append(bc.Blocks, newBlock)
		return
	}
	if isValid(*newBlock, *bc.Blocks[len(bc.Blocks) - 1]) {
		bc.Blocks = append(bc.Blocks, newBlock)
	} else {
		log.Fatal("Warning: Invalid block")
	}

}

// 打印区块链数据
func (bc *Blockchain) Print() {
	for _, block := range bc.Blocks {
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("Previous Hash: %s\n", block.PrevBlockHash)
		fmt.Printf("Current Hash: %s\n", block.Hash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Timestamp: %d\n", block.Timestamp)
		fmt.Println()
	}
}

// new block验证机制
func isValid(newBlock Block, oldBlock Block) bool {
	if (newBlock.Index - 1) != oldBlock.Index {
		return false
	}
	if newBlock.PrevBlockHash != oldBlock.Hash {
		return false
	}
	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}