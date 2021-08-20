//my blockchain program-2
//hashing name+ctime
package main

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type Block struct {
	timestamp    time.Time
	transactions []string
	prevHash     []byte
	Hash         []byte
}

func main() {
	genesisTransactions := []string{"50", "30"}
	var name = "GT"
	genesisBlock := NewBlock(genesisTransactions, []byte{}, name)
	fmt.Println("--- First Block ---" + name)
	printBlockInformation(genesisBlock)

	block2Transactions := []string{"40", "30"}
	name = "B2T"
	block2 := NewBlock(block2Transactions, genesisBlock.Hash, name)
	fmt.Println("--- Second Block ---" + name)
	printBlockInformation(block2)

	block3Transactions := []string{"30", "30"}
	name = "B3T"
	block3 := NewBlock(block3Transactions, block2.Hash, name)
	fmt.Println("--- Third Block ---" + name)
	printBlockInformation(block3)
}

func NewBlock(transactions []string, prevHash []byte, name string) *Block {
	currentTime := time.Now()

	return &Block{
		timestamp:    currentTime,
		transactions: transactions,
		prevHash:     prevHash,
		Hash:         NewHash(currentTime, transactions, prevHash, name),
	}
}

func NewHash(time time.Time, transactions []string, prevHash []byte, name string) []byte {
	input := append(prevHash, time.String()...)
	for transaction := range transactions {
		input = append(input, string(rune(transaction))...)
		input = append(input, name...) //append name
	}

	hash := sha256.Sum256(input)
	return hash[:]
}

func printBlockInformation(block *Block) {
	fmt.Printf("\ttime: %s\n", block.timestamp.String())
	fmt.Printf("\tprevHash: %x\n", block.prevHash)
	fmt.Printf("\tHash: %x\n", block.Hash)
	printTransactions(block)
}

func printTransactions(block *Block) {
	fmt.Println("\tTransactions:")
	for i, transaction := range block.transactions {
		fmt.Printf("\t\t%v: %q\n", i, transaction)
	}
}
