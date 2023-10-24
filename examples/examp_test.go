package examples

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"main/batchtrans"
	"main/sbt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Test_wallet_write(t *testing.T) {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		fmt.Printf("Eth connect error:%s\n", err)
		log.Fatal(err)
	}
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
	//gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)           // in wei
	auth.GasLimit = uint64(80000000)     // in units
	auth.GasPrice = big.NewInt(int64(8)) //gasPrice
	//fmt.Println(fromAddress)

	fornum := 1 // Number of Transfer Address
	to_address := make([]common.Address, 0, fornum)
	value := make([]*big.Int, 0, fornum)

	for i := 0; i < fornum; i++ {
		priveKey, _ := crypto.GenerateKey()
		publicKey := priveKey.Public()
		publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
		address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
		to_address = append(to_address, common.HexToAddress(address))
		value = append(value, big.NewInt(10000)) // Value of Transfer  * 10**14, example 10000 = 1ETH
	}

	scaddress := common.HexToAddress("0x1331eEA5C43c5491BD7aeAA5c6CFefc117EEDa5A") // Smart Contract Address
	instance, err := batchtrans.NewBatchtrans(scaddress, client)
	if err != nil {
		fmt.Println("error creating instance")
		log.Fatal(err)
	}

	tx, err := instance.HandOut(auth, to_address, value)
	if err != nil {
		fmt.Println("error creating tx")
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s\n", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870
	//time.Sleep(10 * time.Second)
	// number, _ := instance.Gettestnum(nil)
	// fmt.Println("new Number: ", number)
}

func Test_SBT(t *testing.T) {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		fmt.Printf("Eth connect error:%s\n", err)
		log.Fatal(err)
	}
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
	//gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)           // in wei
	auth.GasLimit = uint64(80000000)     // in units
	auth.GasPrice = big.NewInt(int64(8)) //gasPrice

	scaddress := common.HexToAddress("0xD0BF8353273cdBA469b1613c214FBBBf6765ecF3") // Smart Contract Address
	instance, err := sbt.NewSbt(scaddress, client)
	if err != nil {
		fmt.Println("error creating instance")
		log.Fatal(err)
	}

	// setTokenInfo
	tokenid := big.NewInt(1)
	totalamount := big.NewInt(5000)
	name := "clayert"
	symbol := "clt"
	url := "https://ipfs.io/ipfs/QmW948aN4Tjh4eLkAAo8os1AcM2FJjA46qtaEfFAnyNYzY"

	tx, err := instance.SettokenIDInfo(auth, tokenid, totalamount, name, symbol, url)
	if err != nil {
		fmt.Println("error creating tx")
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())

}
