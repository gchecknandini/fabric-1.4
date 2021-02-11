package main

import (
	"encoding/json"
	"fmt"

	//"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

//SmartContract ... The SmartContract
type SmartContract struct {
}

//const     Channel = "twoorgschannel"
//const     TargetChaincode = "autonomy_chain"


type Cameras struct {
	Filename string `json:"file_name"`
	CameraId   string `json:"camera_id"`
	IncidentType      string `json:"Incident_type"`
	Latitude string `json:"latitude"`
	Longitude    string `json:"longitude"`
	SavingTimestamp string `json:"saving_timestamp"`
	Timestamp    string `json:"timestamp"`
}

type CameraByIdResponse struct {
	ID      string  `json:"id"`
	Request Cameras `json:"camera"`
}

type Response struct {
	Status  string             `json:"status"`
	Message string             `json:"message"`
	Data    CameraByIdResponse `json:"data"`
}

//var logger = shim.NewLogger("example_cc0")

//Init Function
func (s *SmartContract) Init(stub shim.ChaincodeStubInterface) peer.Response {

	_, args := stub.GetFunctionAndParameters()
	var cameras = Cameras{
		Filename: args[0],
		CameraId:   args[1],
		IncidentType:      args[2],
		Latitude: args[3],
		Longitude:    args[4],
		SavingTimestamp: args[5],
		Timestamp:    args[6]}

	cameraAsBytes, _ := json.Marshal(cameras)

	var uniqueID = args[1]

	err := stub.PutState(uniqueID, cameraAsBytes)

	if err != nil {
		fmt.Println("Error in Init")
	}

	return shim.Success([]byte("Chaincode Successfully initialized"))
}

//Createcamera ... this function is used to create cameras
func CreateCamera(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}
	var cameras = Cameras{
		Filename: args[0],
		CameraId:   args[1],
		IncidentType:      args[2],
		Latitude: args[3],
		Longitude:    args[4],
		SavingTimestamp: args[5],
		Timestamp:    args[6]}

	cameraAsBytes, _ := json.Marshal(cameras)

	var uniqueID = args[1]

	//args := make([][]byte, 1)
	//args[0] = []byte("queryCar")
	//args[1] = []byte[args[1]]

		// Gets the value of MyToken in token chaincode (V5)
	//response := stub.InvokeChaincode(TargetChaincode, args, Channel)
	//if car.statue =	'Active'
	

	err := stub.PutState(uniqueID, cameraAsBytes)

	//else error

	if err != nil {
		fmt.Println("Erro in create camera")
	}

	return shim.Success(nil)
}

//GetCameraByID ... This function will return a particular camera
func GetCameraByID(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	fmt.Println("Before Len")

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments.Expected 1 argument")
	}

	fmt.Println("After Len")

	query := `{
				"selector": {
				   "camera_id": {
					  "$eq": "` + args[0] + `"
				   }
				}
			 }`

	fmt.Println("Queeryy =>>>> \n" + query)

	//resultsIterator, err := stub.GetQueryResult("{\"selector\":{\"doc_type\":\"cameras\",\"_id\":{\"$eq\": \"1\"}}}")

	resultsIterator, err := stub.GetQueryResult(query)

	if err != nil {
		fmt.Println("Error fetching reuslts")
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	fmt.Println("After results")
	// counter := 0
	//var resultKV
	for resultsIterator.HasNext() {
		fmt.Println("Inside hasnext")
		// Get the next element
		queryResponse, err := resultsIterator.Next()

		if err != nil {
			fmt.Println("Err=" + err.Error())
			return shim.Success([]byte("Error in parse: " + err.Error()))
		}

		// Increment the counter
		// counter++
		key := queryResponse.GetKey()
		value := string(queryResponse.GetValue())

		// Print the receieved result on the console
		fmt.Printf("Result#   %s   %s \n\n", key, value)
		b, je := json.Marshal(value)
		if je != nil {
			return shim.Error(je.Error())
		}

		return shim.Success(b)
	}

	// Close the iterator
	resultsIterator.Close()
	return shim.Success(nil)

	//	return shim.Error("Could not find any cameras.")

}

//Invoke function
func (s *SmartContract) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fun, args := stub.GetFunctionAndParameters()
	if fun == "CreateCamera" {
		fmt.Println("Error occured ==> ")
		//logger.Info("########### create docs ###########")
		return CreateCamera(stub, args)
	} else if fun == "GetCameraByID" {
		fmt.Println("Calling get  ==> ")
		return GetCameraByID(stub, args)
	}

	return shim.Error(fmt.Sprintf("Unknown action, check the first argument, must be one of 'delete', 'query', or 'move'. But got: %v", fun))

}

func main() {
	err := shim.Start(new(SmartContract))

	if err != nil {
		fmt.Print(err)
	}
}
