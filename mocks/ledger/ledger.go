package ledger

import "sync"

// ? Struct that will handle the data, mapping the chaincode and the bytes in it
type Db struct {
	Data map[string]map[string][]byte
	sync.RWMutex
}
