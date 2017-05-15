package main
import(
	"fmt";
	"strconv";
	"github.com/hyperledger/fabric/core/chaincode/shim"

)

type VotingChainCode struct{}

func (t *VotingChainCode) Init(stub shim.ChaincodeStubInterface, function string, ar[] string) ([]byte, error){

	var err error
	err = stub.PutState("icecream", []byte("0"));
	if err!=nil {
		return nil, err;
	}
	err = stub.PutState("pizza", []byte("0"));
	if err!=nil {
		return nil, err;
	}
	return nil,nil;
}

func (t *VotingChainCode) Invoke(stub shim.ChaincodeStubInterface, function string, ar[] string) ([]byte, error){
		var err error
		var val int
		var store string
		option,err := stub.GetState(ar[0]);
		if err!=nil{
			return nil, err;
		}
		val, err = strconv.Atoi(string(option))
		store =  strconv.Itoa(val+1);
		err = stub.PutState(ar[0],[]byte(store));
		if err!=nil {
			return nil, err;
		}
		return nil, nil;
}

func (t* VotingChainCode) Query(stub shim.ChaincodeStubInterface, function string, ar[] string) ([]byte, error){
	var err error
	result,err := stub.GetState(ar[0]);
	if err!=nil {
	 	return nil, err;
	 } 
	 return []byte(result), nil;
}

func main(){
	var err error
	err = shim.Start(new(VotingChainCode))
	if err!=nil{
		fmt.Printf("Error starting voting chaincode: %s", err)
	}
}
