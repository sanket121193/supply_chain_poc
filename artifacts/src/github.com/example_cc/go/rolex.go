package main

import (
	"encoding/json"
	"fmt"
	"bytes"
	"strconv"
	"time"
	
//	"strings"
	

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type Product struct {
	SerialNo      	string `json:"serialNo"`
	BatchId       	string `json:"batchId"`
	ModelNo       	string `json:"modelNo"`
	ModelName     	string `json:"modelName"`
	Date     	string `json:"date"`
	CftId         	string `json:"cftId"`
	Price         	string `json:"price"`
	Spec 		string `json:"spec"`
	Status       	string `json:"status"`
	Ownership     	string `json:"ownership"`
	DealerId      	string `json:"dealerId"`
	DealerName    	string `json:"dealerName"`
	Insurance     	string `json:"insurance"`
	InsuranceId     string `json:"insuranceId"`
        InsuranceExpiry string `json:"insuranceExpiry"`
        ServiceId     	string `json:"serviceId"`
	ServiceHistory 	string `json:"serviceHistory"`
	SerCentre	string `json:"serCentre"`
	DateOfService   string `json:"dateOfService"`
	ServiceDetails	string `json:"serviceDetails"`
}

type Service struct {
	
	ServiceId     	string `json:"serviceId"`
	SerCentre   	string `json:"serCentre"`
	DateOfService   string `json:"dateOfService"`
	ServiceDetails	string `json:"serviceDetails"`
}

type AllWatches struct{
	AllWatches []string
}

type AllWatchesDetails struct{
	Watches []Product
}

//type TimeTracker struct{
//	DispachedTime	string
//	ReachedTime	string
//	TPTemprature	string
//}

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {

	fmt.Println("chaincode_custom Init")
	var err error
	var watches AllWatches
	
	watchesAsBytes,_ :=json.Marshal(watches)
	err = stub.PutState("AllWatches", watchesAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("chaincode_custom Invoke")
	function, args := stub.GetFunctionAndParameters()

	if function == "addProduct"{
		return t.addProduct(stub, args)
	} 
	if function == "transferMtoD"{
		return t.transferMtoD(stub,args)
	}
	if function == "transferDtoC"{
		return t.transferDtoC(stub,args)
	}
	if function == "serviceUpdate"{
		return t.serviceUpdate(stub,args)
	}
	if function == "getAllWatches"{
		return t.getAllWatches(stub, args)
	}
	if function == "query" {
		// queries an entity state
		return t.query(stub, args)
	}
	if function == "getproducthistory" {
		// queries an entity state
		fmt.Println("chaincode getproducthistory")
		return t.getproducthistory(stub, args)
	}

	return shim.Error("Invalid invoke function name.")
}

//function to add product

func (t *SimpleChaincode) addProduct(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("Add product...")
	var product Product
	var err error
	
	//isExistAsBytes,_ := stub.GetState(args[0])

	//if(isExistAsBytes != nil){
	//	return shim.Error(err.Error())
	//}

	product.SerialNo=args[0]
	product.BatchId=args[1]
	product.ModelNo=args[2]
	product.ModelName=args[3]
	product.Date=args[4]
	product.CftId=args[5]
	product.Price=args[6]
	product.Spec=args[7]
	product.Status="MFG"
	product.Ownership="Rolex"
	product.DealerId=""
	product.DealerName=""
	product.Insurance="--"
	product.InsuranceId="--"
    	product.InsuranceExpiry="--"
	product.ServiceId="--"
	product.ServiceHistory="--"
	product.SerCentre="--"
	product.DateOfService="--"
	product.ServiceDetails="--"
        
	
	productAsBytes, _ := json.Marshal(product)
	err = stub.PutState(product.SerialNo, productAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

//update all batches to the array whose key is "AllWatches"
	AllWatchesAsBytes, err := stub.GetState("AllWatches")

	var allWatches AllWatches

	err= json.Unmarshal(AllWatchesAsBytes, &allWatches)
	allWatches.AllWatches=append(allWatches.AllWatches,product.SerialNo)

	allwatchesAsBytes,_ :=json.Marshal(allWatches)
	err = stub.PutState("AllWatches", allwatchesAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
//-----------------------------------------------------------
	return shim.Success(nil)
}

//Change the Product by Batch ID
func (t *SimpleChaincode) transferMtoD(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// must be an invoke
	var err error
	var product Product
	bAsBytes, err := stub.GetState(args[0])

	err = json.Unmarshal(bAsBytes, &product)

	if err != nil {
		return shim.Error(err.Error())
	}
	
	product.DealerId=args[1]
	product.DealerName=args[2]
	product.Date = args[3]
	product.Status="Dealer"
	product.Ownership=args[2]
	
	//Commit updates batch to ledger
	btAsBytes, _ := json.Marshal(product)
	err = stub.PutState(product.SerialNo, btAsBytes)	
	if err != nil {
		return shim.Error(err.Error())
	}	
	
	 return shim.Success(nil);
}

func (t *SimpleChaincode) transferDtoC(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// must be an invoke
	var err error
	var product Product
	bAsBytes, err := stub.GetState(args[0])

	err = json.Unmarshal(bAsBytes, &product)

	if err != nil {
		return shim.Error(err.Error())
	}
	  
	product.Date=args[1]
	product.Insurance=args[2]
	product.InsuranceId=args[5]
	product.Status="Customer"
	product.Ownership=args[3]
	product.InsuranceExpiry=args[4]		//expiry date field updated
	
	//Commit updates batch to ledger
	btAsBytes, _ := json.Marshal(product)
	err = stub.PutState(product.SerialNo, btAsBytes)	
	if err != nil {
		return shim.Error(err.Error())
	}	
	
        return shim.Success(nil);
}


func (t *SimpleChaincode) serviceUpdate(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// must be an invoke
	var err error
	var service Service
	var product Product
        productAsBytes, err := stub.GetState(args[0])

	err = json.Unmarshal(productAsBytes, &product)

	if err != nil {
		return shim.Error(err.Error())
	}
	
	product.ServiceId = args[1]
	product.ServiceHistory = args[2]
	product.SerCentre = args[3]
	product.DateOfService = args[4]
	product.ServiceDetails = args[5]

	// must be an invoke
	
	//serviceAsBytes, err := stub.GetState(args[1])

	//err = json.Unmarshal(serviceAsBytes, &service)

	//if err != nil {
	//	return shim.Error(err.Error())
	//}

	service.ServiceId= args[1]
	service.SerCentre = args[3]
	service.DateOfService = args[4]
	service.ServiceDetails = args[5]
	
	
	//Commit updates batch to ledger
	stAsBytes, _ := json.Marshal(service)
	err = stub.PutState( product.SerialNo + service.ServiceId, stAsBytes)	
	if err != nil {
		return shim.Error(err.Error())
	}	
 
	
	//Commit updates batch to ledger
	ptAsBytes, _ := json.Marshal(product)
	err = stub.PutState(product.SerialNo, ptAsBytes)	//ServiceId or SerialNo
	if err != nil {
		return shim.Error(err.Error())
	}
	
	
	return shim.Success(nil);
}

// ============================================================================================================================
// Get All Batches Details for Transporter
// ============================================================================================================================
func (t *SimpleChaincode) getAllWatches(stub shim.ChaincodeStubInterface, args []string) pb.Response{
	
	//get the AllBatches index
	var owner string
	owner =args[0]
	fmt.Printf("Value of Owner: %s", owner)
	allBAsBytes,_ := stub.GetState("AllWatches")
	
	var res AllWatches
	json.Unmarshal(allBAsBytes, &res)
	
	var rab AllWatchesDetails

	for i := range res.AllWatches{

		sbAsBytes,_ := stub.GetState(res.AllWatches[i])
		
		var sb Product
		json.Unmarshal(sbAsBytes, &sb)

	if(sb.Status == owner) {
		fmt.Printf("Value of Owner-1: %s", sb.Status)
		rab.Watches = append(rab.Watches,sb); 
	}
	}
	rabAsBytes, _ := json.Marshal(rab)
	return shim.Success(rabAsBytes)
}


//End of changing the Batch ID

/// Query callback representing the query of a chaincode
func (t *SimpleChaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var serl  string // Entities
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the person to query")
	}

	serl = args[0]

	// Get the state from the ledger
	Avalbytes, err := stub.GetState(serl)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + serl + "\"}"
		return shim.Error(jsonResp)
	}

	if Avalbytes == nil {
		jsonResp := "{\"Error\":\"Nil amount for " + serl + "\"}"
		return shim.Error(jsonResp)
	}

	//	jsonResp := "{\"Name\":\"" + A + "\",\"Amount\":\"" + string(Avalbytes) + "\"}"
	//	logger.Infof("Query Response:%s\n", jsonResp)
	return shim.Success(Avalbytes)
}


//Query to get the history of the BAtchID

func (t *SimpleChaincode) getproducthistory(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Printf("In getproducthistory Function")

	if len(args) < 1 {
		fmt.Printf("In getproducthistory Error Function")
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	SerialNo := args[0]

	fmt.Printf("- start getHistoryForWatch: %s\n", SerialNo)

	resultsIterator, err := stub.GetHistoryForKey(SerialNo)

	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing historic values for the marble
	var buffer bytes.Buffer

	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")
		// if it was a delete operation on given key, then we need to set the
		//corresponding value null. Else, we will write the response.Value
		//as-is (as the Value itself a JSON marble)
		if response.IsDelete {
			buffer.WriteString("null")
		} else {
			buffer.WriteString(string(response.Value))
		}

		buffer.WriteString(", \"Timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		buffer.WriteString("\"")

		buffer.WriteString(", \"IsDelete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(response.IsDelete))
		buffer.WriteString("\"")

		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getHistoryForMarble returning:\n%s\n", buffer.String())

return shim.Success(buffer.Bytes())
}

// ============================================================================================================================


func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
