// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"main/sbt"
	"time"

	//"os"
	"io/ioutil"
	"math/big"

	//"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var(
	tokenid=int64(0);
	totalamount=int64(0);
	mintamount=int64(0);
	receiver = common.HexToAddress("0");
	batchid = make([]int64,0);
	batchamount = make([]int64,0);
	batchreceiver = make([]common.Address,0);
	client *ethclient.Client
)

func buildConn() *ethclient.Client {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		fmt.Printf("Eth connect error:%s\n", err)
		log.Fatal(err)
	}
	return client
}

func getPrivatefromKeystore(ksfile string, pass string) *ecdsa.PrivateKey {
	keyjson, err := ioutil.ReadFile(ksfile) 
	key, err := keystore.DecryptKey(keyjson, pass)
	if err != nil {
		panic(err)
	}
	return key.PrivateKey
}

func getPrivatefromPriKey(prikey string) *ecdsa.PrivateKey {
	privateKey, _ := crypto.HexToECDSA(prikey) 
	return  privateKey
}

func buildTx(prikey string) *bind.TransactOpts {
	client := buildConn()
	defer client.Close()
	privateKey:= getPrivatefromPriKey(prikey)
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
	auth.GasLimit = uint64(30000000) // in units
	auth.GasPrice = gasPrice         //big.NewInt(int64(8))
	return auth
}

func buildTxByDecryKeyStore(ksfile string, pass string) *bind.TransactOpts {
	client := buildConn()
	defer client.Close()
	privateKey:= getPrivatefromKeystore(ksfile, pass)

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
	auth.GasLimit = uint64(30000000) // in units
	auth.GasPrice = gasPrice         //big.NewInt(int64(8))
	return auth
}
func init(){
	client = buildConn()
}


func buildTxByUnlockKeyStore(ksfile string, pass string) *bind.TransactOpts {
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := ioutil.ReadFile(ksfile)
	if err != nil {
	  log.Fatal(err)
	}
	acc, err := ks.Import(jsonBytes, pass, pass)
	if err != nil {
		//fmt.Println("ErrInImport: ",err)
	}
	
	//fmt.Println(acc.Address.Hex())

	err = ks.Unlock(acc,pass)
	if err !=nil{
		//fmt.Println("ErrInUnlock: ",err)
	}

	nonce, err := client.PendingNonceAt(context.Background(), acc.Address)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(nonce)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	chainID,_ := client.ChainID(context.Background())
	//fmt.Println(chainID)
	// auth := bind.NewKeyedTransactor(privateKey)
	auth,_ := bind.NewKeyStoreTransactorWithChainID(ks,acc,chainID)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)       // in wei
	auth.GasLimit = uint64(80000000) // in units
	auth.GasPrice = gasPrice         //big.NewInt(int64(8))
	return auth
}

func setTokeninfo(Txauth *bind.TransactOpts, instance *sbt.Sbt, id int64, amount int64) {
	// setTokenInfo
	tokenid := big.NewInt(id)
	totalamount := big.NewInt(amount)
	name := "clayert"
	symbol := "clt"
	url := "https://ipfs.io/ipfs/QmW948aN4Tjh4eLkAAo8os1AcM2FJjA46qtaEfFAnyNYzY"

	// ids,err:=instance.TokenIDs(nil);
	// if err != nil {
	// 	fmt.Println("error getting initialized tokenIDs: ", err)

	// }
	
	// if len(ids) !=0{
	// 	for _,id :=range ids{
	// 		if id.Int64() == tokenid.Int64(){
	// 			fmt.Println("token already exists")
	// 			return
	// 	    }
	//     }
	// }

	tx, err := instance.SettokenIDInfo(Txauth, tokenid, totalamount, name, symbol, url)
	if err != nil {
		fmt.Println("error creating tx: ", err)
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())
}

func mint(Txauth *bind.TransactOpts, instance *sbt.Sbt,tokenid int64,mintamount int64,receiver common.Address) {
	tx, err := instance.Mint(Txauth, receiver, big.NewInt(tokenid), big.NewInt(mintamount))
	if err != nil {
		fmt.Println("error creating tx")
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())
}

func batchmint(Txauth *bind.TransactOpts, instance *sbt.Sbt,batchid []int64,batchamount []int64,batchreceiver []common.Address) {
	tx, err := instance.Batchmint(Txauth, batchreceiver, tracferIntToBigInt(batchid), tracferIntToBigInt(batchamount))
	if err != nil {
		fmt.Println("error creating tx")
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())
}

func main() {
	// client := buildConn()
	// defer client.Close()
	//prikey := "bebb5b73e288c580a6cee5070929ab3ff8985422d7a0bc45938faae5332e2e2f"
	//Txauth := buildTx(prikey)
	ketstore:="/opt/etherData/keystore/UTC--2023-09-08T03-15-52.105540382Z--596e8070f9b3c607c0d309ed904324844100029a"
	pass:="yu201219jing"
	// Txauth :=buildTxByDecryKeyStore(ketstore,pass)
	Txauth :=buildTxByUnlockKeyStore(ketstore,pass)
	scaddress := common.HexToAddress("0xF2AC45ca3ED21312BB81E72cAe8dc14Cba97214c") // Smart Contract Address
	instance, err := sbt.NewSbt(scaddress, client)
	if err != nil {
		fmt.Println("error creating instance")
		log.Fatal(err)
	}

	tokenid=5;
	mintamount=1;
	receiver=common.HexToAddress("0x2d8Fac7B7295A2aBf75D49A534b3a86920de51B2")
	go mint(Txauth, instance,tokenid,mintamount,receiver)

	// time.Sleep(5*time.Second)
	// Txauth =buildTxByUnlockKeyStore(ketstore,pass)
	// tokenid=2;
	// totalamount=300;
	// setTokeninfo(Txauth, instance,tokenid,totalamount)

	// time.Sleep(5*time.Second)

	// batchTxNum := 1500

	// for i:=0;i<batchTxNum;i++ {
	// 	priveKey, _ := crypto.GenerateKey()
	// 	publicKey := priveKey.Public()
	// 	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	// 	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	// 	batchid = append(batchid, int64(i))
	// 	batchamount = append(batchamount, int64(i))
	// 	batchreceiver = append(batchreceiver, common.HexToAddress(address))
	// }
	batchid = []int64{1,2};
	batchamount = []int64{55,13};
	batchreceiver =[]common.Address{
		common.HexToAddress("0x2d8Fac7B7295A2aBf75D49A534b3a86920de51B2"),
		common.HexToAddress("0x596e8070F9B3C607c0d309ED904324844100029A"),
		//common.HexToAddress("0xe579aBE4a3B4BaB0b8E0791"),
		};
	go batchmint(buildTxByUnlockKeyStore(ketstore,pass), instance,batchid,batchamount,batchreceiver)

	go mint(buildTxByUnlockKeyStore(ketstore,pass), instance,tokenid,mintamount,receiver)
	go mint(buildTxByUnlockKeyStore(ketstore,pass), instance,tokenid,mintamount,receiver)
	go mint(buildTxByUnlockKeyStore(ketstore,pass), instance,tokenid,mintamount,receiver)
	go mint(buildTxByUnlockKeyStore(ketstore,pass), instance,tokenid,mintamount,receiver)
	go mint(buildTxByUnlockKeyStore(ketstore,pass), instance,tokenid,mintamount,receiver)
	go mint(buildTxByUnlockKeyStore(ketstore,pass), instance,tokenid,mintamount,receiver)
	go mint(buildTxByUnlockKeyStore(ketstore,pass), instance,tokenid,mintamount,receiver)
	go mint(buildTxByUnlockKeyStore(ketstore,pass), instance,tokenid,mintamount,receiver)
	go mint(buildTxByUnlockKeyStore(ketstore,pass), instance,tokenid,mintamount,receiver)
	go mint(buildTxByUnlockKeyStore(ketstore,pass), instance,tokenid,mintamount,receiver)
	go mint(buildTxByUnlockKeyStore(ketstore,pass), instance,tokenid,mintamount,receiver)
	go mint(buildTxByUnlockKeyStore(ketstore,pass), instance,tokenid,mintamount,receiver)
	go mint(buildTxByUnlockKeyStore(ketstore,pass), instance,tokenid,mintamount,receiver)
	go mint(buildTxByUnlockKeyStore(ketstore,pass), instance,tokenid,mintamount,receiver)
	go mint(buildTxByUnlockKeyStore(ketstore,pass), instance,tokenid,mintamount,receiver)
	go mint(buildTxByUnlockKeyStore(ketstore,pass), instance,tokenid,mintamount,receiver)

	time.Sleep(5*time.Second)
}

func tracferIntToBigInt(num []int64)(res[]*big.Int){
res = make([]*big.Int, 0)
 for _, v := range num{
res = append(res, big.NewInt(v))
}
return res
}
