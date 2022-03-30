/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/stretchr/testify/mock"
)

//const getStateError = "world state get error"

type MockStub struct {
	shim.ChaincodeStubInterface
	mock.Mock
}

func (ms *MockStub) GetState(key string) ([]byte, error) {
	args := ms.Called(key)

	return args.Get(0).([]byte), args.Error(1)
}

func (ms *MockStub) PutState(key string, value []byte) error {
	args := ms.Called(key, value)

	return args.Error(0)
}

func (ms *MockStub) DelState(key string) error {
	args := ms.Called(key)

	return args.Error(0)
}

type MockContext struct {
	contractapi.TransactionContextInterface
	mock.Mock
}

func (mc *MockContext) GetStub() shim.ChaincodeStubInterface {
	args := mc.Called()

	return args.Get(0).(*MockStub)
}

/*

func configureStub() (*MockContext, *MockStub) {
	var nilBytes []byte

	testWorkshop := new(Workshop)
	testWorkshop.Value = "set value"
	workshopBytes, _ := json.Marshal(testWorkshop)

	ms := new(MockStub)
	ms.On("GetState", "statebad").Return(nilBytes, errors.New(getStateError))
	ms.On("GetState", "missingkey").Return(nilBytes, nil)
	ms.On("GetState", "existingkey").Return([]byte("some value"), nil)
	ms.On("GetState", "workshopkey").Return(workshopBytes, nil)
	ms.On("PutState", mock.AnythingOfType("string"), mock.AnythingOfType("[]uint8")).Return(nil)
	ms.On("DelState", mock.AnythingOfType("string")).Return(nil)

	mc := new(MockContext)
	mc.On("GetStub").Return(ms)

	return mc, ms
}

func TestWorkshopExists(t *testing.T) {
	var exists bool
	var err error

	ctx, _ := configureStub()
	c := new(WorkshopContract)

	exists, err = c.WorkshopExists(ctx, "statebad")
	assert.EqualError(t, err, getStateError)
	assert.False(t, exists, "should return false on error")

	exists, err = c.WorkshopExists(ctx, "missingkey")
	assert.Nil(t, err, "should not return error when can read from world state but no value for key")
	assert.False(t, exists, "should return false when no value for key in world state")

	exists, err = c.WorkshopExists(ctx, "existingkey")
	assert.Nil(t, err, "should not return error when can read from world state and value exists for key")
	assert.True(t, exists, "should return true when value for key in world state")
}

func TestCreateWorkshop(t *testing.T) {
	var err error

	ctx, stub := configureStub()
	c := new(WorkshopContract)

	err = c.CreateWorkshop(ctx, "statebad", "some value")
	assert.EqualError(t, err, fmt.Sprintf("Could not read from world state. %s", getStateError), "should error when exists errors")

	err = c.CreateWorkshop(ctx, "existingkey", "some value")
	assert.EqualError(t, err, "The asset existingkey already exists", "should error when exists returns true")

	err = c.CreateWorkshop(ctx, "missingkey", "some value")
	stub.AssertCalled(t, "PutState", "missingkey", []byte("{\"value\":\"some value\"}"))
}

func TestReadWorkshop(t *testing.T) {
	var workshop *Workshop
	var err error

	ctx, _ := configureStub()
	c := new(WorkshopContract)

	workshop, err = c.ReadWorkshop(ctx, "statebad")
	assert.EqualError(t, err, fmt.Sprintf("Could not read from world state. %s", getStateError), "should error when exists errors when reading")
	assert.Nil(t, workshop, "should not return Workshop when exists errors when reading")

	workshop, err = c.ReadWorkshop(ctx, "missingkey")
	assert.EqualError(t, err, "The asset missingkey does not exist", "should error when exists returns true when reading")
	assert.Nil(t, workshop, "should not return Workshop when key does not exist in world state when reading")

	workshop, err = c.ReadWorkshop(ctx, "existingkey")
	assert.EqualError(t, err, "Could not unmarshal world state data to type Workshop", "should error when data in key is not Workshop")
	assert.Nil(t, workshop, "should not return Workshop when data in key is not of type Workshop")

	workshop, err = c.ReadWorkshop(ctx, "workshopkey")
	expectedWorkshop := new(Workshop)
	expectedWorkshop.Value = "set value"
	assert.Nil(t, err, "should not return error when Workshop exists in world state when reading")
	assert.Equal(t, expectedWorkshop, workshop, "should return deserialized Workshop from world state")
}

func TestUpdateWorkshop(t *testing.T) {
	var err error

	ctx, stub := configureStub()
	c := new(WorkshopContract)

	err = c.UpdateWorkshop(ctx, "statebad", "new value")
	assert.EqualError(t, err, fmt.Sprintf("Could not read from world state. %s", getStateError), "should error when exists errors when updating")

	err = c.UpdateWorkshop(ctx, "missingkey", "new value")
	assert.EqualError(t, err, "The asset missingkey does not exist", "should error when exists returns true when updating")

	err = c.UpdateWorkshop(ctx, "workshopkey", "new value")
	expectedWorkshop := new(Workshop)
	expectedWorkshop.Value = "new value"
	expectedWorkshopBytes, _ := json.Marshal(expectedWorkshop)
	assert.Nil(t, err, "should not return error when Workshop exists in world state when updating")
	stub.AssertCalled(t, "PutState", "workshopkey", expectedWorkshopBytes)
}

func TestDeleteWorkshop(t *testing.T) {
	var err error

	ctx, stub := configureStub()
	c := new(WorkshopContract)

	err = c.DeleteWorkshop(ctx, "statebad")
	assert.EqualError(t, err, fmt.Sprintf("Could not read from world state. %s", getStateError), "should error when exists errors")

	err = c.DeleteWorkshop(ctx, "missingkey")
	assert.EqualError(t, err, "The asset missingkey does not exist", "should error when exists returns true when deleting")

	err = c.DeleteWorkshop(ctx, "workshopkey")
	assert.Nil(t, err, "should not return error when Workshop exists in world state when deleting")
	stub.AssertCalled(t, "DelState", "workshopkey")
}
*/
