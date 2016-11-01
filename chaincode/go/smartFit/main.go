
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
	"errors"
	"strconv"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/op/go-logging"
)