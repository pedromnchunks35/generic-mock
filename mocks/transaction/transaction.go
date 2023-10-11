package transaction

import (
	"sync"

	stub "github.com/pedromnchunks35/generic-mock/mocks/stub"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// ? Get the transaction context interface
type Transaction struct {
	//? the methods it needs to be the transaction
	contractapi.TransactionContext
	//? the stub that we need to give
	Stub *stub.Stub
	//? Create a mutex just for the access of the stub
	StubMutex sync.RWMutex
}

// ? Function to get the stub
func (transaction *Transaction) GetStub() shim.ChaincodeStubInterface {
	//? Handle the locks
	transaction.StubMutex.RLock()
	defer transaction.StubMutex.RUnlock()
	//? return simply the stub
	return transaction.Stub
}
