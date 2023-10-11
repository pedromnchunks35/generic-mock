package test

import (
	"fmt"
	"strings"
	"testing"
)

func Test_Init_Ledger(t *testing.T) {
	asset, err := SmartContract.ReadAsset(TransactionContext, "asset1")
	if err != nil {
		t.Fatalf("it should not throw a error, asset1 is valid")
	}
	if asset.Color != "blue" {
		t.Fatalf("asset1 should be blue")
	}
	if asset.Size != 1 {
		t.Fatalf("asset1 should have a size of 1")
	}
	if asset.Owner != "Pedro Silva" {
		t.Fatalf("asset 1 owner should be Pedro Silva")
	}
	if asset.AppraisedValue != 10 {
		t.Fatalf("the appraised value should be 10")
	}
	fmt.Println("asset1 OK")
	asset, err = SmartContract.ReadAsset(TransactionContext, "asset2")
	if err != nil {
		t.Fatalf("it should not throw a error, asset2 is valid")
	}
	if asset.Color != "red" {
		t.Fatalf("asset2 should be red")
	}
	if asset.Size != 2 {
		t.Fatalf("asset2 should have a size of 2")
	}
	if asset.Owner != "Elsa Silva" {
		t.Fatalf("asset 2 owner should be Elsa Silva")
	}
	if asset.AppraisedValue != 2 {
		t.Fatalf("the appraised value should be 2")
	}
	fmt.Println("asset2 OK")
	asset, err = SmartContract.ReadAsset(TransactionContext, "asset3")
	if err != nil {
		t.Fatalf("it should not throw a error, asset3 is valid")
	}
	if asset.Color != "brown" {
		t.Fatalf("asset3 should be brown")
	}
	if asset.Size != 3 {
		t.Fatalf("asset3 should have a size of 3")
	}
	if asset.Owner != "Bruna Silva" {
		t.Fatalf("asset 3 owner should be Bruna Silva")
	}
	if asset.AppraisedValue != 0 {
		t.Fatalf("the appraised value should be 0")
	}
	fmt.Println("asset3 OK")
}

func Test_Create_Asset_Invalid(t *testing.T) {
	err := SmartContract.CreateAsset(
		TransactionContext,
		"asset3",
		"yellow",
		2,
		"Zé Couve",
		22,
	)
	if !strings.Contains(err.Error(), "the asset asset3 already exists") {
		t.Fatalf("it should throw a error saying that the asset already exists")
	}
	fmt.Println("asset already exists error OK")
}

func Test_Create_Asset(t *testing.T) {
	err := SmartContract.CreateAsset(
		TransactionContext,
		"asset4",
		"yellow",
		2,
		"Zé Couve",
		22,
	)
	if err != nil {
		t.Fatalf("it should not throw a error, asset 4 is valid %v", err)
	}
	asset, err := SmartContract.ReadAsset(TransactionContext, "asset4")
	if err != nil {
		t.Fatalf("it should not throw a error, asset3 is valid")
	}
	if asset.Color != "yellow" {
		t.Fatalf("asset4 should be yellow")
	}
	if asset.Size != 2 {
		t.Fatalf("asset4 should have a size of 2")
	}
	if asset.Owner != "Zé Couve" {
		t.Fatalf("asset 4 owner should be Zé Couve")
	}
	if asset.AppraisedValue != 22 {
		t.Fatalf("the appraised value should be 22")
	}
	fmt.Println("asset4 creation OK")
}

func Test_Read_Asset_Invalid(t *testing.T) {
	_, err := SmartContract.ReadAsset(TransactionContext, "name_that_does_not_exist")
	if !strings.Contains(err.Error(), "the asset name_that_does_not_exist does not exist") {
		t.Fatalf("it should throw a error saying that the asset does not exist")
	}
	fmt.Println("read invalid asset OK")
}

func Test_Update_Asset_Invalid(t *testing.T) {
	err := SmartContract.UpdateAsset(
		TransactionContext,
		"name_that_does_not_exist",
		"color",
		3,
		"Xuyong",
		22,
	)
	if !strings.Contains(err.Error(), "The asset name_that_does_not_exist does not exist") {
		t.Fatalf("should throw a error saying that the id does not exist")
	}
	fmt.Println("update invalid OK")
}

func Test_Update_Asset(t *testing.T) {
	err := SmartContract.UpdateAsset(
		TransactionContext,
		"asset4",
		"javali",
		4,
		"Hugo Barracas",
		7,
	)
	if err != nil {
		t.Fatalf("it should not throw a error, the update is legit for asset4")
	}
	asset, err := SmartContract.ReadAsset(TransactionContext, "asset4")
	if err != nil {
		t.Fatalf("it should not throw a error, asset4 is valid")
	}
	if asset.Color != "javali" {
		t.Fatalf("asset4 should be javali")
	}
	if asset.Size != 4 {
		t.Fatalf("asset4 should have a size of 4")
	}
	if asset.Owner != "Hugo Barracas" {
		t.Fatalf("asset 4 owner should be Hugo Barracas")
	}
	if asset.AppraisedValue != 7 {
		t.Fatalf("the appraised value should be 7")
	}
	fmt.Println("asset4 update OK")
}

func Test_Delete_Asset_Invalid(t *testing.T) {
	err := SmartContract.DeleteAsset(
		TransactionContext,
		"name_that_does_not_exist",
	)
	if !strings.Contains(err.Error(), "The asset name_that_does_not_exist does not exit") {
		t.Fatalf("should throw a error saying that the asset does not exist")
	}
	fmt.Println("invalid asset delete OK")
}

func Test_Delete_Asset(t *testing.T) {
	err := SmartContract.DeleteAsset(
		TransactionContext,
		"asset4",
	)
	if err != nil {
		t.Fatalf("should not throw a erro, delete the asset4 is legit")
	}
	_, err = SmartContract.ReadAsset(
		TransactionContext,
		"asset4",
	)
	if !strings.Contains(err.Error(), "the asset asset4 does not exist") {
		t.Fatalf("should throw a error saying that the asset4 does not exist")
	}
	fmt.Println("delete asset4 OK")
}

func Test_Transfer_Asset_Invalid(t *testing.T) {
	err := SmartContract.TransferAsset(
		TransactionContext,
		"name_that_does_not_exist",
		"Kublau",
	)
	if !strings.Contains(err.Error(), "the asset name_that_does_not_exist does not exist") {
		t.Fatalf("should throw a error saying that the asset does not exist")
	}
	fmt.Println("invalid transfer OK")
}

func Test_Transfer_Asset(t *testing.T) {
	err := SmartContract.TransferAsset(
		TransactionContext,
		"asset3",
		"Kublau",
	)
	if err != nil {
		t.Fatalf("should not throw a error since it is a valid transfer")
	}
	asset, err := SmartContract.ReadAsset(TransactionContext, "asset3")
	if err != nil {
		t.Fatalf("should not throw a error since asset3 is valid")
	}
	if asset.Owner != "Kublau" {
		t.Fatalf("it should have the new owner name, since it got updated")
	}
	fmt.Println("transfer OK")
}

func Test_Get_All_Assets(t *testing.T) {
	assets, err := SmartContract.GetAllAssets(TransactionContext)
	if err != nil {
		t.Fatalf("should not throw a error %v", err)
	}
	//? Check first asset
	if assets[0].Color != "blue" {
		t.Fatalf("asset1 should be blue")
	}
	if assets[0].Size != 1 {
		t.Fatalf("asset1 should have a size of 1")
	}
	if assets[0].Owner != "Pedro Silva" {
		t.Fatalf("asset 1 owner should be Pedro Silva")
	}
	if assets[0].AppraisedValue != 10 {
		t.Fatalf("the appraised value should be 10")
	}
	fmt.Println("asset1 OK")
	//? Check Secound asset
	if assets[1].Color != "red" {
		t.Fatalf("asset2 should be red")
	}
	if assets[1].Size != 2 {
		t.Fatalf("asset2 should have a size of 2")
	}
	if assets[1].Owner != "Elsa Silva" {
		t.Fatalf("asset 2 owner should be Elsa Silva")
	}
	if assets[1].AppraisedValue != 2 {
		t.Fatalf("the appraised value should be 2")
	}
	fmt.Println("asset2 OK")
	//? Check third asset
	if assets[2].Color != "brown" {
		t.Fatalf("asset3 should be brown")
	}
	if assets[2].Size != 3 {
		t.Fatalf("asset3 should have a size of 3")
	}
	if assets[2].Owner != "Kublau" {
		t.Fatalf("asset 3 owner should be Kublau")
	}
	if assets[2].AppraisedValue != 0 {
		t.Fatalf("the appraised value should be 0")
	}
	fmt.Println("asset3 OK")
}
