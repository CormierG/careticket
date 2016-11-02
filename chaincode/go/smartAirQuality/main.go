
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


// SmartAirQualityChaincode example Chaincode implementation
type SmartAirQualityChaincode struct {

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


type SmartAirQualityState struct {
    AssetID          string      `json:"assetid,omitempty"`        // all assets must have an ID, primary key of contract
    Temperature      string      `json:"temperature,omitempty"`    // -40 to +125°C (-40 to +257°F) / +- 0.3°C
    Humidity         string      `json:"humidity,omitempty"`       // 0 to 95% / +- 3%
    CO2              string      `json:"co2,omitempty"`            // 0 to 4000 ppm / +- 75ppm
    Dust             string      `json:"dust,omitempty"`           // 0 to 500 μg/m3
    VOCs             string      `json:"vocs,omitempty"`
    Timestamp        string      `json:"timestamp,omitempty"`        
}