// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"main/sbt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func buildConn() *ethclient.Client {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		fmt.Printf("Eth connect error:%s\n", err)
		log.Fatal(err)
	}
	return client
}

func buildTx() *bind.TransactOpts {
	client := buildConn()
	defer client.Close()
	privateKey, _ := crypto.HexToECDSA("bebb5b73e288c580a6cee5070929ab3ff8985422d7a0bc45938faae5332e2e2f")
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)       // in wei
	auth.GasLimit = uint64(80000000) // in units
	auth.GasPrice = gasPrice         //big.NewInt(int64(8))
	return auth
}

func setTokeninfo(Txauth *bind.TransactOpts, instance *sbt.Sbt) {
	// setTokenInfo
	tokenid := big.NewInt(1)
	totalamount := big.NewInt(5000)
	name := "clayert"
	symbol := "clt"
	url := "https://ipfs.io/ipfs/QmW948aN4Tjh4eLkAAo8os1AcM2FJjA46qtaEfFAnyNYzY"

	tx, err := instance.SettokenIDInfo(Txauth, tokenid, totalamount, name, symbol, url)
	if err != nil {
		fmt.Println("error creating tx")
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())
}

func mint(Txauth *bind.TransactOpts, instance *sbt.Sbt) {
	// setTokenInfo
	account := common.HexToAddress("")
	id := big.NewInt(1)
	amount := big.NewInt(20)

	tx, err := instance.Mint(Txauth, account, id, amount)
	if err != nil {
		fmt.Println("error creating tx")
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())
}

func batchmint(Txauth *bind.TransactOpts, instance *sbt.Sbt) {
	// setTokenInfo
	account := []common.Address{common.HexToAddress("")}
	id := []*big.Int{big.NewInt(1)}
	amount := []*big.Int{big.NewInt(20)}

	tx, err := instance.Batchmint(Txauth, account, id, amount)
	if err != nil {
		fmt.Println("error creating tx")
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())
}

func main() {
	client := buildConn()
	defer client.Close()
	Txauth := buildTx()
	scaddress := common.HexToAddress("0xD0BF8353273cdBA469b1613c214FBBBf6765ecF3") // Smart Contract Address
	instance, err := sbt.NewSbt(scaddress, client)
	if err != nil {
		fmt.Println("error creating instance")
		log.Fatal(err)
	}
	setTokeninfo(Txauth, instance)
}
