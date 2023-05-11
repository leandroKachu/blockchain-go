package main

import (
	"os"

	"github.com/tensor-programming/golang-blockchain/wallet"
)

func main() {
	defer os.Exit(0)
	// cli := cli.CommandLine{}
	// cli.Run()
	w := wallet.MakeWallet()

	w.Address()

}
