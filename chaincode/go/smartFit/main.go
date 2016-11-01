
/*
Contributors:

Mike Aro - Initial Contribution
Jay Venenga - Initial Contribution
Bryan Kribbs - Initial Contribution


November 2016
*/

package main

import (
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
    "time"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)


// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {

}

const CONTRACTSTATEKEY string = "ContractStateKey"  
// store contract state - only version in this example
const MYVERSION string = "1.0"

// ************************************
// contract state 
// ************************************


type ContractState struct {
    Version      string                        `json:"version"`
}


type SmartFitState struct {
    AssetID          string      `json:"assetid,omitempty"`        // all assets must have an ID, primary key of contract
    ActiveMinutes    string      `json:"activeminutes,omitempty"`       
    CaloriesOut      string      `json:"caloriesout,omitempty"`    
    Steps            string      `json:"steps,omitempty"`
    Timestamp        string      `json:"timestamp,omitempty"`        
}