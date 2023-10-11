package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// ? STRUCT FOR MANAGING A ASSET
type SmartContract struct {
	contractapi.Contract
}

// ? Struct for the asset itself, denote that we tell how we want the json to be formated
type Asset struct {
	AppraisedValue int    `json:"AppraisedValue"`
	Color          string `json:"Color"`
	ID             string `json:"ID"`
	Owner          string `json:"Owner"`
	Size           int    `json:"Size"`
}

/*
* @receiver s, the struct of the smart contract iself
* @param ctx, the interface for the api of the contract, which we will use to invoke methods against the stub
* @returns nothing or error.When nothing it inits the ledger, otherwise it throws a error
 */
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	//? Create some assets
	assets := []Asset{
		{ID: "asset1", Color: "blue", Size: 1, Owner: "Pedro Silva", AppraisedValue: 10},
		{ID: "asset2", Color: "red", Size: 2, Owner: "Elsa Silva", AppraisedValue: 2},
		{ID: "asset3", Color: "brown", Size: 3, Owner: "Bruna Silva", AppraisedValue: 0},
	}
	//? Loop all over the assets and convert them to json
	for _, asset := range assets {
		//? Convert asset to json
		assetJson, err := json.Marshal(asset)
		if err != nil {
			return err
		}
		//? Put the asset into the ledger
		err = ctx.GetStub().PutState(asset.ID, assetJson)
		if err != nil {
			return fmt.Errorf("Failed to put to world state %v", err)
		}
	}
	return nil
}

/*
* @receiver s, the struct of the smart contract iself
* @param ctx, the interface for the api of the contract, which we will use to invoke methods against the stub
* @param id, the id of the asset
* @param color, the color we whish our new asset to have
* @param size, the size of the asset we want to store
* @param owner, the owner of the asset
* @param appraisedValue, the appraised value for the new asset
* @returns nothing or error. Case of nothing the asset is created, otherwise it throws a error
 */
func (s *SmartContract) CreateAsset(
	ctx contractapi.TransactionContextInterface,
	id string,
	color string,
	size int,
	owner string,
	appraisedValue int,
) error {
	// ? Check if exists
	exists, err := s.AssetExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the asset %s already exists", id)
	}
	//? Create asset
	asset := Asset{
		ID:             id,
		Color:          color,
		Size:           size,
		Owner:          owner,
		AppraisedValue: appraisedValue,
	}
	//? Marshall asset
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return fmt.Errorf("error encoding the asset %v", err)
	}
	return ctx.GetStub().PutState(id, assetJSON)
}

/*
* @receiver s, the struct of the SmartContract
* @param ctx, the interface of the contract api
* @param id, the id which we will use to query the asset
* @returns *Asset or error. *Asset, the pointer of the asset if it is success, otherwise it throws a error
 */
func (s *SmartContract) ReadAsset(
	ctx contractapi.TransactionContextInterface,
	id string,
) (*Asset, error) {
	//? Get asset and check err
	assetJson, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state: %v", err)
	}
	if assetJson == nil {
		return nil, fmt.Errorf("the asset %s does not exist", id)
	}
	//? Unmarshal the asset and return it if possible
	var asset Asset
	err = json.Unmarshal(assetJson, &asset)
	if err != nil {
		return nil, err
	}
	return &asset, nil
}

/*
* @receiver s, the struct of the SmartContrat
* @param ctx, the interface for the contract api
* @param id, the asset we want to overwrite
* @param color, the color we whish our new asset to have
* @param size, the size of the asset we want to store
* @param owner, the owner of the asset
* @param appraisedValue, the appraised value for the new asset
* @returns nothing or error. Case nothing it means success on overwriting, otherwise it throws a error
 */
func (s *SmartContract) UpdateAsset(
	ctx contractapi.TransactionContextInterface,
	id string,
	color string,
	size int,
	owner string,
	appraisedValue int,
) error {
	//? Check if the asset exists
	exists, err := s.AssetExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("The asset %s does not exist", id)
	}
	//? Create the asset that will overwrite the one with the same key and then overwrite it
	asset := Asset{
		ID:             id,
		Color:          color,
		Size:           size,
		Owner:          owner,
		AppraisedValue: appraisedValue,
	}
	assetJson, err := json.Marshal(asset)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(id, assetJson)
}

/*
* @receiver s, the struct of the SmartContrat
* @param ctx, the interface for the contract api
* @param id, the asset we want to delete
* @returns nothing or error. Case nothing delete got success, otherwise it will throw a error
 */
func (s *SmartContract) DeleteAsset(
	ctx contractapi.TransactionContextInterface,
	id string,
) error {
	//? Check if asset exists,case it does return it
	exists, err := s.AssetExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("The asset %s does not exit", id)
	}
	return ctx.GetStub().DelState(id)
}

/*
* @receiver s, the struct of the Smart Contract
* @param ctx, the interface for the contract api
* @param id, the id of the asset we whish to check existance of
* @returns bool or error. Case the asset exists, true. Case not false and error
 */
func (s *SmartContract) AssetExists(
	ctx contractapi.TransactionContextInterface,
	id string,
) (bool, error) {
	assetJson, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("Failed to read from world state: %v", err)
	}
	return assetJson != nil, nil
}

/*
* @receiver s, the struct of the Smart Contract
* @param ctx, the interface for the contract api
* @param id, the id of the asset we whish to check existance of
* @param newOwner, the new owner for the asset
* @returns nothing or error. Nothing case success changing the owner, error case it throws a erro
 */
func (s *SmartContract) TransferAsset(
	ctx contractapi.TransactionContextInterface,
	id string,
	newOwner string,
) error {
	asset, err := s.ReadAsset(ctx, id)
	if err != nil {
		return err
	}
	asset.Owner = newOwner
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(id, assetJSON)
}

/*
* @receiver s, the struct of the Smart Contract
* @param ctx, the interface for the contract api
* @returns []*Asset or error. []*Asset, array of pointers case success, otherwise a error
 */
func (s *SmartContract) GetAllAssets(ctx contractapi.TransactionContextInterface) ([]*Asset, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	//? Close after all the tasks get done
	defer resultsIterator.Close()
	var assets []*Asset
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		var asset Asset
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return nil, err
		}
		assets = append(assets, &asset)
	}
	return assets, nil
}
