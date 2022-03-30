/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-contract-api-go/metadata"
)

func main() {
	workshopContract := new(WorkshopContract)
	workshopContract.Info.Version = "0.0.1"
	workshopContract.Info.Description = "My Smart Contract"
	workshopContract.Info.License = new(metadata.LicenseMetadata)
	workshopContract.Info.License.Name = "Apache-2.0"
	workshopContract.Info.Contact = new(metadata.ContactMetadata)
	workshopContract.Info.Contact.Name = "John Doe"

	chaincode, err := contractapi.NewChaincode(workshopContract)
	chaincode.Info.Title = "fabric_e-workshop chaincode"
	chaincode.Info.Version = "0.0.1"

	if err != nil {
		panic("Could not create chaincode from WorkshopContract." + err.Error())
	}

	err = chaincode.Start()

	if err != nil {
		panic("Failed to start chaincode. " + err.Error())
	}
}
