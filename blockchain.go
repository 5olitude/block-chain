package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	transactions []string
}

func NewBlock(nonce int, previousHash string) *Block {
	b := new(Block)
	b.nonce = nonce
	b.previousHash = previousHash
	b.timestamp = time.Now().UnixNano()
	return b
}

func (b *Block) Print() {
	fmt.Printf("timestamp%d\n", b.timestamp)
	fmt.Printf("nonce%d\n", b.nonce)
	fmt.Printf("previous_hash%s\n", b.previousHash)
	fmt.Printf("transactions%s\n", b.transactions)
}

type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

func NewBlockchain() *Blockchain {
	bc := new(Blockchain)
	bc.chain = []*Block{}          // Initialize chain as empty slice of Block pointers
	bc.CreateBlock(0, "init hash") // Create genesis block
	return bc
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash string) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b) // Append new block to chain
	return b
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d  %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 25))
}

func init() {
	log.SetPrefix("BlockChain :")
}

func main() {
	blockChain := NewBlockchain()
	blockChain.Print()

	// Add blocks to the blockchain
	blockChain.CreateBlock(5, "hash 1")
	blockChain.Print()

	blockChain.CreateBlock(2, "hash 2")
	blockChain.Print()
}
