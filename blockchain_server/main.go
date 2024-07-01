package main

import (
	"flag"
	"log"
)

func init() {
	log.SetPrefix("BlockChain :")
}

func main() {
	port := flag.Uint("port", 5001, "TCP Port Number for blockchain server")
	flag.Parse()
	app := NewBlockchainServer(uint16(*port))
	app.Run()
}
