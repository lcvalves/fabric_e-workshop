/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// WorkshopContract contract for managing CRUD for Lot
type WorkshopContract struct {
	contractapi.Contract
}

/*
 ***************************
 ***************************
 * LOT TRANSCATION METHDOS *
 ***************************
 ***************************
 */

// LotExists returns true when lot with given ID exists in world state
func (c *WorkshopContract) LotExists(ctx contractapi.TransactionContextInterface, lotID string) (bool, error) {
	data, err := ctx.GetStub().GetState(lotID)

	if err != nil {
		return false, err
	}

	return data != nil, nil
}

// CreateLot creates a new instance of Lot
func (c *WorkshopContract) CreateLot(ctx contractapi.TransactionContextInterface, lotID, product string, amount float32,
	unit, owner string) (string, error) {
	exists, err := c.LotExists(ctx, lotID)
	if err != nil {
		return "", fmt.Errorf("could not read from world state. %s", err)
	} else if exists {
		return "", fmt.Errorf("the lot %s already exists", lotID)
	}

	// Validação de quantidade (não deve ser possível criar lotes com nenhuma quantidade)
	if amount < 0 {
		return "", fmt.Errorf("the amount should be greater than 0")
	}

	lot := &Lot{
		DocType: "lot",
		ID:      lotID,
		Product: product,
		Amount:  amount,
		Unit:    unit,
		Owner:   owner,
	}

	bytes, err := json.Marshal(lot)
	if err != nil {
		return "", err
	}

	err = ctx.GetStub().PutState(lot.ID, bytes)
	if err != nil {
		return "", fmt.Errorf("failed to put to world state: %v", err)
	}

	return fmt.Sprintf("%s created successfully", lotID), nil
}

// ReadLot retrieves an instance of Lot from the world state
func (c *WorkshopContract) ReadLot(ctx contractapi.TransactionContextInterface, lotID string) (*Lot, error) {
	exists, err := c.LotExists(ctx, lotID)
	if err != nil {
		return nil, fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return nil, fmt.Errorf("the lot %s does not exist", lotID)
	}

	bytes, _ := ctx.GetStub().GetState(lotID)

	lot := new(Lot)

	err = json.Unmarshal(bytes, lot)

	if err != nil {
		return nil, fmt.Errorf("could not unmarshal world state data to type Lot")
	}

	return lot, nil
}

// constructQueryResponseFromIterator constructs a slice of lots from the resultsIterator
func constructQueryResponseFromIterator(resultsIterator shim.StateQueryIteratorInterface) ([]*Lot, error) {
	var lots []*Lot
	for resultsIterator.HasNext() {
		queryResult, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		var lot Lot
		err = json.Unmarshal(queryResult.Value, &lot)
		if err != nil {
			return nil, err
		}
		lots = append(lots, &lot)
	}

	return lots, nil
}

// getQueryResultForQueryString executes the passed in query string.
// The result set is built and returned as a byte array containing the JSON results.
func getQueryResultForQueryString(ctx contractapi.TransactionContextInterface, queryString string) ([]*Lot, error) {
	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	return constructQueryResponseFromIterator(resultsIterator)
}

// GetAllLots queries for all lots.
// This is an example of a parameterized query where the query logic is baked into the chaincode,
// and accepting a single query parameter (docType).
// Only available on state databases that support rich query (e.g. CouchDB)
// Example: Parameterized rich query
func (c *WorkshopContract) GetAllLots(ctx contractapi.TransactionContextInterface) ([]*Lot, error) {
	queryString := `{"selector":{"docType":"lot"}}`
	return getQueryResultForQueryString(ctx, queryString)
}

// UpdateLot retrieves an instance of Lot from the world state and updates its value
func (c *WorkshopContract) UpdateLot(ctx contractapi.TransactionContextInterface, lotID string, newValue string) error {
	exists, err := c.LotExists(ctx, lotID)
	if err != nil {
		return fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return fmt.Errorf("the lot %s does not exist", lotID)
	}

	lot := new(Lot)
	//lot.Value = newValue

	bytes, _ := json.Marshal(lot)

	return ctx.GetStub().PutState(lotID, bytes)
}

// DeleteLot deletes an instance of Lot from the world state
func (c *WorkshopContract) DeleteLot(ctx contractapi.TransactionContextInterface, lotID string) error {
	exists, err := c.LotExists(ctx, lotID)
	if err != nil {
		return fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return fmt.Errorf("the lot %s does not exist", lotID)
	}

	return ctx.GetStub().DelState(lotID)
}

// DeleteAllLots deletes all lots found in world state
func (c *WorkshopContract) DeleteAllLots(ctx contractapi.TransactionContextInterface) (string, error) {

	lots, err := c.GetAllLots(ctx)
	if err != nil {
		return "", fmt.Errorf("could not read from world state. %s", err)
	} else if len(lots) == 0 {
		return "", fmt.Errorf("there are no lots in world state to delete")
	}

	for _, lot := range lots {
		err = ctx.GetStub().DelState(lot.ID)
		if err != nil {
			return "", fmt.Errorf("could not delete from world state. %s", err)
		}
	}

	return "all the lots were successfully deleted", nil
}

/*
 ********************************
 ********************************
 * ACTIVITY TRANSCATION METHDOS *
 ********************************
 ********************************
 */
/*
// ActivityExists returns true when activity with given ID exists in world state
func (c *WorkshopContract) ActivityExists(ctx contractapi.TransactionContextInterface, activityID string) (bool, error) {
	data, err := ctx.GetStub().GetState(activityID)

	if err != nil {
		return false, err
	}

	return data != nil, nil
}

// CreateActivity creates a new instance of Activity
func (c *WorkshopContract) CreateActivity(ctx contractapi.TransactionContextInterface, activityID string, value string) error {
	exists, err := c.ActivityExists(ctx, activityID)
	if err != nil {
		return fmt.Errorf("could not read from world state. %s", err)
	} else if exists {
		return fmt.Errorf("the activity %s already exists", activityID)
	}

	activity := new(Activity)
	//activity.Value = value

	bytes, _ := json.Marshal(activity)

	return ctx.GetStub().PutState(activityID, bytes)
}

// ReadActivity retrieves an instance of Activity from the world state
func (c *WorkshopContract) ReadActivity(ctx contractapi.TransactionContextInterface, activityID string) (*Activity, error) {
	exists, err := c.ActivityExists(ctx, activityID)
	if err != nil {
		return nil, fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return nil, fmt.Errorf("the activity %s does not exist", activityID)
	}

	bytes, _ := ctx.GetStub().GetState(activityID)

	activity := new(Activity)

	err = json.Unmarshal(bytes, activity)

	if err != nil {
		return nil, fmt.Errorf("could not unmarshal world state data to type Activity")
	}

	return activity, nil
}

// UpdateActivity retrieves an instance of Activity from the world state and updates its value
func (c *WorkshopContract) UpdateActivity(ctx contractapi.TransactionContextInterface, activityID string, newValue string) error {
	exists, err := c.ActivityExists(ctx, activityID)
	if err != nil {
		return fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return fmt.Errorf("the activity %s does not exist", activityID)
	}

	activity := new(Activity)
	//activity.Value = newValue

	bytes, _ := json.Marshal(activity)

	return ctx.GetStub().PutState(activityID, bytes)
}

// DeleteActivity deletes an instance of Activity from the world state
func (c *WorkshopContract) DeleteActivity(ctx contractapi.TransactionContextInterface, activityID string) error {
	exists, err := c.ActivityExists(ctx, activityID)
	if err != nil {
		return fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return fmt.Errorf("the activity %s does not exist", activityID)
	}

	return ctx.GetStub().DelState(activityID)
} */
