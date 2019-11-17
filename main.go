package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/somecookie/Lauzhack2019/contract"
	"log"
	"os"
)

var myenv map[string]string

const envLoc = ".env"
const ErrTransactionWait = "if you've just started the application, wait a while for the network to confirm your transaction."

func loadEnv() {
	var err error
	if myenv, err = godotenv.Read(envLoc); err != nil {
		log.Printf("could not load env from %s: %v", envLoc, err)
	}
}

func NewSession(ctx context.Context) (session contract.ContractSession) {
	loadEnv()
	keystore, err := os.Open(myenv["KEYSTORE"])
	if err != nil {
		log.Printf(
			"could not load keystore from location %s: %v\n",
			myenv["KEYSTORE"],
			err,
		)
	}
	defer keystore.Close()

	keystorepass := myenv["KEYSTOREPASS"]
	auth, err := bind.NewTransactor(keystore, keystorepass)
	if err != nil {
		log.Printf("%s\n", err)
	}

	// Return session without contract instance
	return contract.ContractSession{
		TransactOpts: *auth,
		CallOpts: bind.CallOpts{
			From:    auth.From,
			Context: ctx,
		},
	}
}

// NewContract deploys a contract if no existing contract exists
func NewContract(session contract.ContractSession, client *ethclient.Client) contract.ContractSession {
	loadEnv()

	// Hash answer before sending it over Ethereum network.
	contractAddress, tx, instance, err := contract.DeployContract(&session.TransactOpts, client)
	if err != nil {
		log.Fatalf("could not deploy contract: %v\n", err)
	}
	fmt.Printf("Contract deployed! Wait for tx %s to be confirmed.\n", tx.Hash().Hex())

	session.Contract = instance
	updateEnvFile("CONTRACTADDR", contractAddress.Hex())
	return session
}

// LoadContract loads a contract if one exists
func LoadContract(session contract.ContractSession, client *ethclient.Client) contract.ContractSession {
	loadEnv()

	addr := common.HexToAddress(myenv["CONTRACTADDR"])
	instance, err := contract.NewContract(addr, client)

	if err != nil {
		log.Println(ErrTransactionWait)
		log.Fatalf("could not load contract: %v\n", err)
	}
	session.Contract = instance
	return session
}

// Utility functions
// stringToKeccak256 converts a string to a keccak256 hash of type [32]byte
func stringToKeccak256(s string) [32]byte {
	var output [32]byte
	copy(output[:], crypto.Keccak256([]byte(s))[:])
	return output
}

// updateEnvFile updates our env file with a key-value pair
func updateEnvFile(k string, val string) {
	myenv[k] = val
	err := godotenv.Write(myenv, envLoc)
	if err != nil {
		log.Printf("failed to update %s: %v\n", envLoc, err)
	}
}

func main() {
	loadEnv()

	//ctx := context.Background()

	client, err := ethclient.Dial(myenv["GATEWAY"])
	if err != nil {
		log.Fatalf("could not connect to Ethereum gateway: %v\n", err)
	}
	defer client.Close()

	session := NewSession(context.Background())

	// Load or Deploy contract, and update session with contract instance
	if myenv["CONTRACTADDR"] == "" {
		session = NewContract(session, client)
	}

	// If we have an existing contract, load it; if we've deployed a new contract, attempt to load it.
	if myenv["CONTRACTADDR"] != "" {
		session = LoadContract(session, client)
	}
}
