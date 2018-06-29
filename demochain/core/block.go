package core

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

// 定义一个Block结构体
type Block struct {
	Index int64 // 区块编号
	Timestamp int64 // 区块时间戳
	PrevBlockHash string // 上一个区块Hash值
	Hash string // 当前区块Hash值

	Data string // 区块数据
}

func calculateHash(b Block) string {
	// 生成blockData
	blockData := string(b.Index) + string(b.Timestamp) + b.PrevBlockHash + b.Data

	// 用sha256计算Hash值
	hashInBytes := sha256.Sum256([]byte(blockData))

	return hex.EncodeToString(hashInBytes[:])
}

func GenerateNewBlock(preBlock Block, data string) Block {
	newBlock := Block{}
	newBlock.Index = preBlock.Index + 1
	newBlock.PrevBlockHash = preBlock.Hash
	newBlock.Timestamp = time.Now().Unix()
	newBlock.Data = data
	newBlock.Hash = calculateHash(newBlock)
	return newBlock
}

// 创始区块
func GenerateGenesisBlock() Block {
	// 定义一个创始区块的preBlock
	preBlock := Block{}
	preBlock.Index = -1
	preBlock.Hash = ""
	return GenerateNewBlock(preBlock, "Genesis Block")
}
