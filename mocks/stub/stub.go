package stub

import (
	"fmt"
	ledger "generic-mock/mocks/ledger"
	"sort"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/ledger/queryresult"
)

// ? This is the chaincode itself
type Stub struct {
	shim.ChaincodeStubInterface
	Db *ledger.Db
}

// ? Iterator will implement "StateQueryIteratorInterface"
// ? and also implement the "CommonIteratorInterface", which is the interface that
// ? is inside of "StateQueryIteratorInterface"
// ? this way we can override the methods of both using only the "Iterator" struct
type Iterator struct {
	shim.StateQueryIteratorInterface
	shim.CommonIteratorInterface
	IteratorCashe []string
	Stub          *Stub
}

func (iterator *Iterator) Next() (*queryresult.KV, error) {
	//? Get first member of the cashe
	key := iterator.IteratorCashe[0]
	//? Throw error case we reached the end
	if key == "" {
		return nil, fmt.Errorf("there is no next value")
	}
	//? Assign results
	result := &queryresult.KV{}
	result.Key = key
	result.Value = iterator.Stub.Db.Data["basic"][key]
	result.Namespace = "basic"
	//? Move the cashe
	iterator.IteratorCashe = iterator.IteratorCashe[1:]
	//? return the result
	return result, nil
}

func (iterator *Iterator) HasNext() bool {
	//? Case there is no more items, return false
	if len(iterator.IteratorCashe) == 0 {
		return false
	} else {
		//? case there is more items return true
		return true
	}
}

func (iterator *Iterator) Close() error {
	//? If it is already closed throw error
	if len(iterator.IteratorCashe) == 0 {
		return fmt.Errorf("the cashe was already closed")
	}
	//? Close the cache
	iterator.IteratorCashe = []string{""}
	return nil
}

func (stub *Stub) PutState(key string, value []byte) error {
	//? Locks handling
	stub.Db.RWMutex.Lock()
	defer stub.Db.Unlock()
	//? Check chaincode
	if stub.Db.Data["basic"] == nil {
		return fmt.Errorf("the chaincode was not initiated or does not exist")
	}
	stub.Db.Data["basic"][key] = value
	return nil
}

func (stub Stub) GetState(key string) ([]byte, error) {
	//? Locks handling
	stub.Db.RWMutex.Lock()
	defer stub.Db.Unlock()
	//? Check chaincode
	if stub.Db.Data["basic"] == nil {
		return nil, fmt.Errorf("the chaincode was not initiated or does not exist")
	}
	return stub.Db.Data["basic"][key], nil
}

func (stub Stub) DelState(key string) error {
	//? Locks handling
	stub.Db.RWMutex.Lock()
	defer stub.Db.Unlock()
	//? Check chaincode
	if stub.Db.Data["basic"] == nil {
		return fmt.Errorf("the chaincode was not initiated or does not exist")
	}
	//? Check if value exists
	if stub.Db.Data["basic"][key] == nil {
		return fmt.Errorf("there is nothing to delete")
	}
	//? Make the delete from the map
	delete(stub.Db.Data["basic"], key)
	return nil
}

// ? This function is to get a iterator from the ledger that is filtered by startKey and endKey, which represent
// ? a lexical interval
func (stub Stub) GetStateByRange(startKey string, endKey string) (shim.StateQueryIteratorInterface, error) {
	//? Check if both keys are valid
	_, err := stub.GetState(startKey)
	if err != nil {
		return nil, fmt.Errorf("startKey is not valid")
	}
	_, err = stub.GetState(endKey)
	if err != nil {
		return nil, fmt.Errorf("endKey is not valid")
	}
	//? Creathe the iterator
	result := &Iterator{}
	result.Stub = &stub
	result.IteratorCashe = []string{}
	//? Extract keys
	keys := make([]string, 0, len(stub.Db.Data["basic"]))
	for k := range stub.Db.Data["basic"] {
		keys = append(keys, k)
	}
	//? Sort them
	sort.Strings(keys)
	//? Case the range is unbounded, the cashe will have every key
	if startKey == "" && endKey == "" {
		result.IteratorCashe = keys
	} else if startKey == "" && endKey != "" {
		//? Case the startKey is unbonded but the end is bounded
		endIndex := sort.SearchStrings(keys, endKey)
		result.IteratorCashe = keys[:endIndex]
	} else if startKey != "" && endKey == "" {
		//? Case the startKey is bounded and the endKey is unbounded
		startIndex := sort.SearchStrings(keys, startKey)
		result.IteratorCashe = keys[startIndex:]
	} else {
		//? Case are both bounded
		endIndex := sort.SearchStrings(keys, endKey)
		startIndex := sort.SearchStrings(keys, startKey)
		result.IteratorCashe = keys[startIndex:endIndex]
	}
	return result, nil
}
