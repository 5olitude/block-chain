package main

import (
	"block-chain/wallet"
	"fmt"
	"log"
)

func init() {
	log.SetPrefix("BlockChain :")
}

func main() {
	w := wallet.NewWallet()
	fmt.Println(w.PrivateKeyStr())
	fmt.Println(w.PublicKeyStr())
	fmt.Println(w.BlockchainAddress())
}
