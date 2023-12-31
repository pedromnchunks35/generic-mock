package test

import (
	"log"
	"testing"

	c "github.com/pedromnchunks35/generic-mock/chaincode"
	mocksLedger "github.com/pedromnchunks35/generic-mock/mocks/ledger"
	mocksStub "github.com/pedromnchunks35/generic-mock/mocks/stub"
	mocksTransaction "github.com/pedromnchunks35/generic-mock/mocks/transaction"
)

var TransactionContext *mocksTransaction.Transaction
var SmartContract *c.SmartContract

func TestMain(t *testing.M) {
	//? Init smart contract
	SmartContract := &c.SmartContract{}
	//? Init the ledger
	ledger := &mocksLedger.Db{}
	ledger.Data = make(map[string]map[string][]byte)
	ledger.Data["basic"] = make(map[string][]byte) //? Notice that this is the name of our current chaincode
	//? Init the sub
	stub := &mocksStub.Stub{}
	stub.Db = ledger
	//? Init the transaction context
	TransactionContext = &mocksTransaction.Transaction{}
	TransactionContext.Stub = stub
	//? Init the ledger
	//? In order to init the ledger, we need to have:
	//? Fake stub (which as a function that gets the fuck stub)
	//? Fake ledger for the putstate
	err := SmartContract.InitLedger(TransactionContext)
	if err != nil {
		log.Fatalf("something went wrong with the init of the ledger %v", err)
	}
	t.Run()
}
