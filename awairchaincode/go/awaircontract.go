/*******************************************************************************

Mike Aro

******************************************************************************/

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// AwairChaincode implementation
type AwairChaincode struct {

}

// CONTRACTSTATEKEY  is used to store contract state into world state
const CONTRACTSTATEKEY string = "ContractStatekey"

// MYVERSION must use this to deploy contract
const MYVERSION string = "1.5"


// ************************************
// asset and contract state
// ************************************

// ContractState holds the contract version
type ContractState struct {
	Version string `json:"version"`
}

// Geolocation stores lat and long
type Geolocation struct {
    Latitude  *float64 `json:"latitude,omitempty"`
    Longitude *float64 `json:"longitude,omitempty"`
}

// AssetState stores current state for any asset
type AssetState struct {
	AssetID		*string 	    `json:assetID,omitempty"`
	Co2        	*float64	    `json:co2,omitempty"`
	Dust    	*float64	    `json:dust,omitempty"`
	Temp 		*float64        `json:temp,omitempty"`
	Humid  		*float64	    `json:humid,omitempty"`
	Voc 		*float64	    `json:voc,omitempty"`
    Location    *Geolocation    `json:"location,omitempty"`    // current asset location
    Alert       *string         `json:alert,omitempty"`
    AlertTime   *string         `json:alerttime,omitempty"`      
}


var contractState = ContractState{MYVERSION}

// ************************************
// deploy callback mode
// ************************************

// Init is called during deploy
func (t *AwairChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
    var stateArg ContractState
    var err error
    if len(args) != 1 {
        return nil, errors.New("init expects one argument, a JSON string with tagged version string")
    }
    err = json.Unmarshal([]byte(args[0]), &stateArg)
    if err != nil {
        return nil, errors.New("Version argument unmarshal failed: " + fmt.Sprint(err))
    }
    if stateArg.Version != MYVERSION {
        return nil, errors.New("Contract version " + MYVERSION + " must match version argument: " + stateArg.Version)
    }
    contractStateJSON, err := json.Marshal(stateArg)
    if err != nil {
        return nil, errors.New("Marshal failed for contract state" + fmt.Sprint(err))
    }
    err = stub.PutState(CONTRACTSTATEKEY, contractStateJSON)
    if err != nil {
        return nil, errors.New("Contract state failed PUT to ledger: " + fmt.Sprint(err))
    }
    return nil, nil
}

// ************************************
// deploy and invoke callback mode
// ************************************

// Invoke is called when an invoke message is received
func (t *AwairChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	// Handle different functions
	if function == "createAsset" {
		// create assetID
		return t.createAsset(stub, args)
	} else if function == "updateAsset" {
		// create assetID
		return t.updateAsset(stub, args)
	} else if function == "deleteAsset" {
		// Deletes an asset by ID from the ledger
		return t.deleteAsset(stub, args)
	}
	return nil, errors.New("Received unknown invocation: " + function)
}


// ************************************
// query callback mode
// ************************************

// Query is called when a query message is received
func (t *AwairChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	// Handle different functions
	if function == "readAsset" {
		// gets the state for an assetID as a JSON struct
		return t.readAsset(stub, args)
	} else if function == "readAssetObjectModel" {
		return t.readAssetObjectModel(stub, args)
	} else if function == "readAssetSamples" {
		// returns selected sample objects
		return t.readAssetSamples(stub, args)
	} else if function == "readAssetSchemas" {
		// returns selected sample objects
		return t.readAssetSchemas(stub, args)
	}
	return nil, errors.New("Received unknown invocation: " + function)
}

/**********main implementation *************/
func main() {
	err := shim.Start(new(AwairChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple Chaincode: %s", err)
	}
}


/*****************ASSET CRUD INTERFACE starts here************/

/****************** 'deploy' methods *****************/

/******************** createAsset ********************/

func (t *AwairChaincode) createAsset(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	_, erval := t.createOrUpdateAsset(stub, args)
	return nil, erval
}

//******************** updateAsset ********************/

func (t *AwairChaincode) updateAsset(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    _, erval := t.createOrUpdateAsset(stub, args)
    return nil, erval
}

//******************** deleteAsset ********************/

func (t *AwairChaincode) deleteAsset(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    var assetID string // asset ID
    var err error
    var stateIn AssetState

    // validate input data for number of args, Unmarshaling to asset state and obtain asset id
    stateIn, err = t.validateInput(args)
    if err != nil {
        return nil, err
    }
    assetID = *stateIn.AssetID
    // Delete the key / asset from the ledger
    err = stub.DelState(assetID)
    if err != nil {
        err = errors.New("DELSTATE failed! : " + fmt.Sprint(err))
        return nil, err
    }
    return nil, nil
}

/******************* Query Methods ***************/

//********************readAsset********************/

func (t *AwairChaincode) readAsset(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    var assetID string // asset ID
    var err error
    var state AssetState

    // validate input data for number of args, Unmarshaling to asset state and obtain asset id
    stateIn, err := t.validateInput(args)
    if err != nil {
        return nil, errors.New("Asset does not exist!")
    }
    assetID = *stateIn.AssetID
    // Get the state from the ledger
    assetBytes, err := stub.GetState(assetID)
    if err != nil || len(assetBytes) == 0 {
        err = errors.New("Unable to get asset state from ledger")
        return nil, err
    }
    err = json.Unmarshal(assetBytes, &state)
    if err != nil {
        err = errors.New("Unable to unmarshal state data obtained from ledger")
        return nil, err
    }
    return assetBytes, nil
}

//*************readAssetObjectModel*****************/

func (t *AwairChaincode) readAssetObjectModel(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    var state AssetState = AssetState{}

    // Marshal and return
    stateJSON, err := json.Marshal(state)
    if err != nil {
        return nil, err
    }
    return stateJSON, nil
}

//*************readAssetSamples*******************/

func (t *AwairChaincode) readAssetSamples(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    return []byte(samples), nil
}

//*************readAssetSchemas*******************/

func (t *AwairChaincode) readAssetSchemas(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    return []byte(schemas), nil
}

// ************************************
// validate input data : common method called by the CRUD functions
// ************************************
func (t *AwairChaincode) validateInput(args []string) (stateIn AssetState, err error) {
    var assetID string                  // asset ID
    var state = AssetState{} // The calling function is expecting an object of type AssetState

    if len(args) != 1 {
        err = errors.New("Incorrect number of arguments. Expecting a JSON strings with mandatory assetID")
        return state, err
    }
    jsonData := args[0]
    assetID = ""
    stateJSON := []byte(jsonData)
    err = json.Unmarshal(stateJSON, &stateIn)
    if err != nil {
        err = errors.New("Unable to unmarshal input JSON data")
        return state, err
        // state is an empty instance of asset state
    }
    // was assetID present?
    // The nil check is required because the asset id is a pointer.
    // If no value comes in from the json input string, the values are set to nil

    if stateIn.AssetID != nil {
        assetID = strings.TrimSpace(*stateIn.AssetID)
        if assetID == "" {
            err = errors.New("AssetID not passed")
            return state, err
        }
    } else {
        err = errors.New("Asset id is mandatory in the input JSON data")
        return state, err
    }

    stateIn.AssetID = &assetID
    return stateIn, nil
}

//******************** createOrUpdateAsset ********************/

func (t *AwairChaincode) createOrUpdateAsset(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    var assetID string // asset ID                    // used when looking in map
    var err error
    var stateIn AssetState
    var stateStub AssetState

    // validate input data for number of args, Unmarshaling to asset state and obtain asset id

    stateIn, err = t.validateInput(args)
    if err != nil {
        return nil, err
    }
    assetID = *stateIn.AssetID
    // Partial updates introduced here
    // Check if asset record existed in stub
    assetBytes, err := stub.GetState(assetID)
    if err != nil || len(assetBytes) == 0 {
        // This implies that this is a 'create' scenario
        stateStub = stateIn // The record that goes into the stub is the one that cme in
    } else {
        // This is an update scenario
        err = json.Unmarshal(assetBytes, &stateStub)
        if err != nil {
            err = errors.New("Unable to unmarshal JSON data from stub")
            return nil, err
            // state is an empty instance of asset state
        }
        // Merge partial state updates
        stateStub, err = t.mergePartialState(stateStub, stateIn)
        if err != nil {
            err = errors.New("Unable to merge state")
            return nil, err
        }
    }
    stateJSON, err := json.Marshal(stateStub)
    if err != nil {
        return nil, errors.New("Marshal failed for contract state" + fmt.Sprint(err))
    }
    // Get existing state from the stub

    // Write the new state to the ledger
    err = stub.PutState(assetID, stateJSON)
    if err != nil {
        err = errors.New("PUT ledger state failed: " + fmt.Sprint(err))
        return nil, err
    }
    return nil, nil
}


/*********************************  internal: mergePartialState ****************************/
func (t *AwairChaincode) mergePartialState(oldState AssetState, newState AssetState) (AssetState, error) {

    old := reflect.ValueOf(&oldState).Elem()
    new := reflect.ValueOf(&newState).Elem()
    for i := 0; i < old.NumField(); i++ {
        oldOne := old.Field(i)
        newOne := new.Field(i)
        if !reflect.ValueOf(newOne.Interface()).IsNil() {
            oldOne.Set(reflect.Value(newOne))
        }
    }
    return oldState, nil
}



