package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

var A, B string
var Aval, Bval, X int

// Init callback representing the invocation of a chaincode
// This chaincode will manage two accounts A and B and will transfer X units from A to B upon invoke
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	var err error

	if len(args) != 4 {
		return nil, errors.New("Incorrect number of arguments. Expecting 4")
	}

	// Initialize the chaincode
	A = args[0]
	Aval, err = strconv.Atoi(args[1])
	if err != nil {
		return nil, errors.New("Expecting integer value for asset holding")
	}
	B = args[2]
	Bval, err = strconv.Atoi(args[3])
	if err != nil {
		return nil, errors.New("Expecting integer value for asset holding")
	}
	fmt.Printf("Aval = %d, Bval = %d\n", Aval, Bval)
	return nil, nil
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	// Transaction makes payment of X units from A to B
	var err error
	X, err = strconv.Atoi(args[0])
	Aval = Aval - X
	Bval = Bval + X
	ts, err2 := stub.GetTxTimestamp()
	if err2 != nil {
		fmt.Printf("Error getting transaction timestamp: %s", err2)
	}
	fmt.Printf("Transaction Time: %v,Aval = %d, Bval = %d\n", ts, Aval, Bval)
	return nil, err
}

// Query callback representing the query of a chaincode
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	return nil, nil
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
