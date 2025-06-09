package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// Estrutura de um bloco
type Block struct {
	Index        int
	Timestamp    string
	Data         string
	PrevHash     string
	Hash         string
}

// Função para calcular o hash de um bloco
func calculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp + block.Data + block.PrevHash
	hash := sha256.New()
	hash.Write([]byte(record))
	return hex.EncodeToString(hash.Sum(nil))
}

// Função para criar o primeiro bloco (gênese)
func createGenesisBlock() Block {
	return Block{
		Index:     0,
		Timestamp: time.Now().String(),
		Data:      "Genesis Block",
		PrevHash:  "",
		Hash:      "",
	}
}

// Função para criar um novo bloco
func createNewBlock(prevBlock Block, data string) Block {
	newBlock := Block{
		Index:     prevBlock.Index + 1,
		Timestamp: time.Now().String(),
		Data:      data,
		PrevHash:  prevBlock.Hash,
		Hash:      "",
	}
	// Calcula o hash do novo bloco
	newBlock.Hash = calculateHash(newBlock)
	return newBlock
}

// Função para criar e imprimir a blockchain com loop infinito
func main() {
	// Criar o bloco gênese
	genesisBlock := createGenesisBlock()

	// Criar a blockchain e adicionar o bloco gênese
	blockchain := []Block{genesisBlock}

	// Imprimir o bloco gênese imediatamente
	fmt.Printf("Index: %d\n", genesisBlock.Index)
	fmt.Printf("Timestamp: %s\n", genesisBlock.Timestamp)
	fmt.Printf("Data: %s\n", genesisBlock.Data)
	fmt.Printf("Previous Hash: %s\n", genesisBlock.PrevHash)
	fmt.Printf("Hash: %s\n", genesisBlock.Hash)
	fmt.Println()

	// Loop infinito para criar blocos a cada 3 segundos
	for {
		// Cria um novo bloco com dados fictícios
		newBlock := createNewBlock(blockchain[len(blockchain)-1], fmt.Sprintf("Block Data %d", blockchain[len(blockchain)-1].Index+1))

		// Adiciona o novo bloco à blockchain
		blockchain = append(blockchain, newBlock)

		// Imprime o novo bloco
		fmt.Printf("Index: %d\n", newBlock.Index)
		fmt.Printf("Timestamp: %s\n", newBlock.Timestamp)
		fmt.Printf("Data: %s\n", newBlock.Data)
		fmt.Printf("Previous Hash: %s\n", newBlock.PrevHash)
		fmt.Printf("Hash: %s\n", newBlock.Hash)
		fmt.Println()

		// Aguarda 3 segundos antes de criar o próximo bloco
		time.Sleep(2 * time.Second)
	}
}
