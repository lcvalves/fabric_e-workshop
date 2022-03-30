/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// WorkshopContract contract for managing CRUD for Workshop
type WorkshopContract struct {
	contractapi.Contract
}

// WorkshopExists returns true when asset with given ID exists in world state
func (c *WorkshopContract) WorkshopExists(ctx contractapi.TransactionContextInterface, workshopID string) (bool, error) {
	data, err := ctx.GetStub().GetState(workshopID)

	if err != nil {
		return false, err
	}

	return data != nil, nil
}

// CreateWorkshop creates a new instance of Workshop
func (c *WorkshopContract) CreateWorkshop(ctx contractapi.TransactionContextInterface, workshopID string, value string) error {
	exists, err := c.WorkshopExists(ctx, workshopID)
	if err != nil {
		return fmt.Errorf("Could not read from world state. %s", err)
	} else if exists {
		return fmt.Errorf("The asset %s already exists", workshopID)
	}

	workshop := new(Workshop)
	workshop.Value = value

	bytes, _ := json.Marshal(workshop)

	return ctx.GetStub().PutState(workshopID, bytes)
}

// ReadWorkshop retrieves an instance of Workshop from the world state
func (c *WorkshopContract) ReadWorkshop(ctx contractapi.TransactionContextInterface, workshopID string) (*Workshop, error) {
	exists, err := c.WorkshopExists(ctx, workshopID)
	if err != nil {
		return nil, fmt.Errorf("Could not read from world state. %s", err)
	} else if !exists {
		return nil, fmt.Errorf("The asset %s does not exist", workshopID)
	}

	bytes, _ := ctx.GetStub().GetState(workshopID)

	workshop := new(Workshop)

	err = json.Unmarshal(bytes, workshop)

	if err != nil {
		return nil, fmt.Errorf("Could not unmarshal world state data to type Workshop")
	}

	return workshop, nil
}

// UpdateWorkshop retrieves an instance of Workshop from the world state and updates its value
func (c *WorkshopContract) UpdateWorkshop(ctx contractapi.TransactionContextInterface, workshopID string, newValue string) error {
	exists, err := c.WorkshopExists(ctx, workshopID)
	if err != nil {
		return fmt.Errorf("Could not read from world state. %s", err)
	} else if !exists {
		return fmt.Errorf("The asset %s does not exist", workshopID)
	}

	workshop := new(Workshop)
	workshop.Value = newValue

	bytes, _ := json.Marshal(workshop)

	return ctx.GetStub().PutState(workshopID, bytes)
}

// DeleteWorkshop deletes an instance of Workshop from the world state
func (c *WorkshopContract) DeleteWorkshop(ctx contractapi.TransactionContextInterface, workshopID string) error {
	exists, err := c.WorkshopExists(ctx, workshopID)
	if err != nil {
		return fmt.Errorf("Could not read from world state. %s", err)
	} else if !exists {
		return fmt.Errorf("The asset %s does not exist", workshopID)
	}

	return ctx.GetStub().DelState(workshopID)
}
