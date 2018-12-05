/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

//WARNING - this chaincode's ID is hard-coded in chaincode_example04 to illustrate one way of
//calling chaincode from a chaincode. If this example is modified, chaincode_example04.go has
//to be modified as well with the new ID of chaincode_example02.
//chaincode_example05 show's how chaincode ID can be passed in as a parameter instead of
//hard-coding.

import (
	"fmt"
	"strconv"
	"encoding/gob"
    "bytes"
    "encoding/json"
    "strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

type Transaction struct{
	Sender string `json:"sender"`
	Receiver string `json:"receiver"`
	TxAmount string `json:"txamount"`
	PaymentMedium string `json:"paymentmedium"`
	TxNo string `json:"txno"`
	//var Date time  
}

func GetBytes(key interface{}) ([]byte, error) {
    var buf bytes.Buffer
    enc := gob.NewEncoder(&buf)
    err := enc.Encode(key)
    if err != nil {
        return nil, err
    }
    return buf.Bytes(), nil
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("ex02 Init")
	_, args := stub.GetFunctionAndParameters()
	//var Txmap, TxValidmap string    // Entities
	//var Aval, Bval int // Asset holdings
	var TxArray map[string]Transaction
	var TxValidArray map[string]string
	var lastID string
	var err error
	TxArray = make(map[string]Transaction)
	TxValidArray = make(map[string]string)
	var lastIDval = 900788600000

	if len(args) != 4 {
		//return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	// Initialize the chaincode
/*	A = args[0]
	Aval, err = strconv.Atoi(args[1])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}
	B = args[2]
	Bval, err = strconv.Atoi(args[3])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}
	fmt.Printf("Aval = %d, Bval = %d\n", Aval, Bval)

	// Write the state to the ledger
	err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(B, []byte(strconv.Itoa(Bval)))
	if err != nil {
		return shim.Error(err.Error())
	}  */
	var TxArrayval []byte
	TxArrayval,err = json.Marshal(TxArray)
	err = stub.PutState("Txmap",[]byte(TxArrayval))
	if err != nil {
		return shim.Error(err.Error())
	}
	var TxValidArrayval []byte
	TxValidArrayval,err = json.Marshal(TxValidArray)
	err = stub.PutState("TxValidmap",[]byte(TxValidArrayval))
	if err != nil {
		return shim.Error(err.Error())
	}
	lastID = strconv.Itoa(lastIDval)
	err = stub.PutState("lastID",[]byte(lastID))
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("ex02 Invoke")
	function, args := stub.GetFunctionAndParameters()
/*	if function == "invoke" {
		// Make payment of X units from A to B
		return t.invoke(stub, args)
	} else if function == "delete" {
		// Deletes an entity from its state
		return t.delete(stub, args)
	} else if function == "query" {
		// the old "Query" is now implemtned in invoke
		return t.query(stub, args) */
	if function == "submit" {
		// Deletes an entity from its state
		return t.submit(stub, args)
	}else if function == "verify" {
		// Deletes an entity from its state
		return t.verify(stub, args)
	}else if function == "view" {
		// Deletes an entity from its state
		return t.view(stub, args)
	}else if function == "mytransactions"{
		//
		return t.mytransactions(stub,args)
	}else if function == "alltransactions"{
		//
		return t.alltransactions(stub,args);
	}
	return shim.Error("Invalid invoke function name. Expecting \"invoke\" \"delete\" \"query\"")
}


/*
// Transaction makes payment of X units from A to B
func (t *SimpleChaincode) invoke(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//var A, B string    // Entities
	//var Aval, Bval int // Asset holdings
	//var X int          // Transaction value
	var err error

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	A = args[0]
	B = args[1]

	// Get the state from the ledger
	// TODO: will be nice to have a GetAllState call to ledger
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if Avalbytes == nil {
		return shim.Error("Entity not found")
	}
	Aval, _ = strconv.Atoi(string(Avalbytes))

	Bvalbytes, err := stub.GetState(B)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if Bvalbytes == nil {
		return shim.Error("Entity not found")
	}
	Bval, _ = strconv.Atoi(string(Bvalbytes))

	// Perform the execution
	X, err = strconv.Atoi(args[2])
	if err != nil {
		return shim.Error("Invalid transaction amount, expecting a integer value")
	}
	Aval = Aval - X
	Bval = Bval + X
	fmt.Printf("Aval = %d, Bval = %d\n", Aval, Bval)

	// Write the state back to the ledger
	err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(B, []byte(strconv.Itoa(Bval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
} */

func (t *SimpleChaincode) submit(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 6 {
		//return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	var Sender = args[0]
	var Receiver = args[1]
	var TxAmount = args[2]
	var PaymentMedium = args[3]
	var TxNo = args[4]
	//var Date = args[5]  
	//var Txmap, TxValidmap,lastID string
	var lastIDval int
	var tx = Transaction{Sender,Receiver,TxAmount,PaymentMedium,TxNo}
	TxArray, err := stub.GetState("Txmap")
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if TxArray == nil {
		return shim.Error("Txmap Entity not found")
	}
	//var TxArrayval interface{} 
	//err = json.Unmarshal(TxArray,&TxArrayval)
	//var TxArrayType = TxArrayval.(map[string]Transaction)
	var TxNewmap map[string]Transaction
	TxNewmap = make(map[string]Transaction)
	var s = strings.Split(string(TxArray),"&")
	//fmt.Println(s)
	for i:=0 ; i<len(s)/6.;i++{
		TxNewmap[s[6*i]] = Transaction{s[6*i+1],s[6*i+2],s[6*i+3],s[6*i+4],s[6*i+5]}
	}
	/*for k,v := range TxArrayType{
		//println(k,v)
		TxNewmap[k] = v
	} */
	TxValidArray, err := stub.GetState("TxValidmap")
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if TxValidArray == nil {
		return shim.Error("TxValidmap Entity not found")
	}
	//var TxValidArrayval interface{}
	//err = json.Unmarshal(TxValidArray,&TxValidArrayval)
	//var TxValidType = TxValidArrayval.(map[string]string)
	var TxValidNewmap map[string]string
	TxValidNewmap = make(map[string]string)
	s = strings.Split(string(TxValidArray),"&")
	for i:=0; i< len(s)/2;i++{
		TxValidNewmap[s[2*i]] = s[2*i+1]
	}
	/*for k,v := range TxValidType{
		TxValidNewmap[k] = v
	} */
	lastIDBytes, err := stub.GetState("lastID")
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if lastIDBytes == nil {
		return shim.Error("lastID Entity not found")
	}
	lastIDval, _ = strconv.Atoi(string(lastIDBytes))
	lastIDval++

	//if (TxValidArray[lastIDval] == 0){
		TxNewmap[string(lastIDBytes)] = tx
		TxValidNewmap[string(lastIDBytes)] = "0"
	//}
	var TxArraystr string
	var buffer bytes.Buffer
	for k,v := range TxNewmap{
		buffer.WriteString(k)
		buffer.WriteString("&")
		buffer.WriteString(v.Sender)
		buffer.WriteString("&")
		buffer.WriteString(v.Receiver)
		buffer.WriteString("&")
		buffer.WriteString(v.TxAmount)
		buffer.WriteString("&")
		buffer.WriteString(v.PaymentMedium)
		buffer.WriteString("&")
		buffer.WriteString(v.TxNo)
		buffer.WriteString("&")
	}
	//TxArrayvalBytes,_ = json.Marshal(TxNewmap)
	TxArraystr = buffer.String()
	err = stub.PutState("Txmap", []byte(TxArraystr))
	if err != nil {
		return shim.Error(err.Error())
	}

	var TxValidArraystr string
	buffer.Reset()
	for k,v := range TxValidNewmap{
		buffer.WriteString(k)
		buffer.WriteString("&")
		buffer.WriteString(v)
		buffer.WriteString("&")
	}
	//TxValidArrayvalBytes,_ = json.Marshal(TxValidNewmap)
	TxValidArraystr = buffer.String()
	err = stub.PutState("TxValidmap", []byte(TxValidArraystr))
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState("lastID", []byte(strconv.Itoa(lastIDval)))
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(lastIDBytes)
}

func (t *SimpleChaincode) verify(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 6 {
		//return shim.Error("Incorrect number of arguments. Expecting 6")
	}
	var requestID = args[0]
	var state = args[1]
	TxValidArray, err := stub.GetState("TxValidmap")
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if TxValidArray == nil {
		return shim.Error("TxValidmap Entity not found")
	}
	//var TxValidArrayval interface{}
	//err = json.Unmarshal(TxValidArray,&TxValidArrayval)
	//var TxValidType = TxValidArrayval.(map[string]string)
	var TxValidNewmap map[string]string
	TxValidNewmap = make(map[string]string)
	var s = strings.Split(string(TxValidArray),"&")
	for i:=0; i< len(s)/2;i++{
		TxValidNewmap[s[2*i]] = s[2*i+1]
	}

	if state == "true" || state == "yes"{
		TxValidNewmap[requestID] = "1"
	}else if state == "false" || state == "no"{
		TxValidNewmap[requestID] = "2"
	}

	var TxValidArraystr string
	var buffer bytes.Buffer
	for k,v := range TxValidNewmap{
		buffer.WriteString(k)
		buffer.WriteString("&")
		buffer.WriteString(v)
		buffer.WriteString("&")
	}
	//TxValidArrayvalBytes,_ = json.Marshal(TxValidNewmap)
	TxValidArraystr = buffer.String()
	err = stub.PutState("TxValidmap", []byte(TxValidArraystr))
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success([]byte("Verification Successful"))
}


func (t *SimpleChaincode) view(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 6 {
		//return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	var requestID = args[0]
	//var Txmap,TxValidmap string
	TxValidArray, err := stub.GetState("TxValidmap")
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if TxValidArray == nil {
		return shim.Error("TxValidmap Entity not found")
	}
	//var TxValidArrayval interface{}
	//err = json.Unmarshal(TxValidArray,&TxValidArrayval)
	//var TxValidType = TxValidArrayval.(map[string]string)
	var TxValidNewmap map[string]string
	TxValidNewmap = make(map[string]string)
	var s = strings.Split(string(TxValidArray),"&")
	for i:=0; i< len(s)/2;i++{
		TxValidNewmap[s[2*i]] = s[2*i+1]
	}
	var state string 
	//var requestIDval int
	//requestIDval,_ = strconv.Atoi(string(requestID))
	state = TxValidNewmap[requestID]
	if (err != nil){
		return shim.Error("Entity not found")
	}
		
	if state=="1" {
		return shim.Success([]byte("Transaction is verified to be True"))
	}else if state== "2"{
		return shim.Success([]byte("Transaction is verified not to be True"))
	}
	return shim.Success([]byte("Transaction is not verified"))
}

func (t *SimpleChaincode) mytransactions(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 6 {
		//return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	var user_name = args[0]
	TxArray, err := stub.GetState("Txmap")
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if TxArray == nil {
		return shim.Error("Txmap Entity not found")
	}
	//var TxArrayval interface{} 
	//err = json.Unmarshal(TxArray,&TxArrayval)
	//var TxArrayType = TxArrayval.(map[string]Transaction)
	var TxNewmap map[string]Transaction
	TxNewmap = make(map[string]Transaction)
	var s = strings.Split(string(TxArray),"&")
	fmt.Println(s)
	for i:=0 ; i<len(s)/6.;i++{
		TxNewmap[s[6*i]] = Transaction{s[6*i+1],s[6*i+2],s[6*i+3],s[6*i+4],s[6*i+5]}
	}
	/*for k,v := range TxArrayType{
		//println(k,v)
		TxNewmap[k] = v
	} */
	TxValidArray, err := stub.GetState("TxValidmap")
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if TxValidArray == nil {
		return shim.Error("TxValidmap Entity not found")
	}

	var TxValidNewmap map[string]string
	TxValidNewmap = make(map[string]string)
	s = strings.Split(string(TxValidArray),"&")
	for i:=0; i< len(s)/2;i++{
		TxValidNewmap[s[2*i]] = s[2*i+1]
	}
	var TxArraystr string
	var keys []string
	var buffer bytes.Buffer
	for k,v := range TxNewmap{
		if v.Sender == user_name || v.Receiver == user_name{
			fmt.Println(k, " - ",v)
			buffer.WriteString(k)
			buffer.WriteString("&")
			buffer.WriteString(v.Sender)
			buffer.WriteString("&")
			buffer.WriteString(v.Receiver)
			buffer.WriteString("&")
			buffer.WriteString(v.TxAmount)
			buffer.WriteString("&")
			buffer.WriteString(v.PaymentMedium)
			buffer.WriteString("&")
			buffer.WriteString(v.TxNo)
			buffer.WriteString("&")
			keys = append(keys,k)
		}
	}
	
	buffer.WriteString("&$$$$$$$&")
	var isPresent bool
	for k,v := range TxValidNewmap{
		isPresent = false
		for i := range keys{
			if (k == keys[i]){
				isPresent = true
				break
			}	
		}
		if isPresent == true{
			buffer.WriteString(k)
			buffer.WriteString("&")
			buffer.WriteString(v)
			buffer.WriteString("&")
		}
	}
	TxArraystr = buffer.String() 
	return shim.Success([]byte(TxArraystr))
}


func (t *SimpleChaincode) alltransactions(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 6 {
		//return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	var lastIDval int
	//var tx = Transaction{Sender,Receiver,TxAmount,PaymentMedium,TxNo}
	TxArray, err := stub.GetState("Txmap")
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if TxArray == nil {
		return shim.Error("Entity not found")
	}
	//var TxArrayval interface{} 
	//err = json.Unmarshal(TxArray,&TxArrayval)
	//var TxArrayType = TxArrayval.(map[string]Transaction)
	var TxNewmap map[string]Transaction
	TxNewmap = make(map[string]Transaction)
	var s = strings.Split(string(TxArray),"&")
	for i:=0 ; i<len(s)/6.;i++{
		TxNewmap[s[6*i]] = Transaction{s[6*i+1],s[6*i+2],s[6*i+3],s[6*i+4],s[6*i+5]}
	}
	/*for k,v := range TxArrayType{
		//println(k,v)
		TxNewmap[k] = v
	} */
	TxValidArray, err := stub.GetState("TxValidmap")
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if TxValidArray == nil {
		return shim.Error("Entity not found")
	}
	//var TxValidArrayval interface{}
	//err = json.Unmarshal(TxValidArray,&TxValidArrayval)
	//var TxValidType = TxValidArrayval.(map[string]string)
	var TxValidNewmap map[string]string
	TxValidNewmap = make(map[string]string)
	s = strings.Split(string(TxValidArray),"&")
	for i:=0; i< len(s)/2;i++{
		TxValidNewmap[s[2*i]] = s[2*i+1]
	}
	/*for k,v := range TxValidType{
		TxValidNewmap[k] = v
	} */
	lastIDBytes, err := stub.GetState("lastID")
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if lastIDBytes == nil {
		return shim.Error("Entity not found")
	}
	lastIDval, _ = strconv.Atoi(string(lastIDBytes))

	var TxArraystr string
	var buffer bytes.Buffer
	for k,v := range TxNewmap{
		buffer.WriteString(k)
		buffer.WriteString("&")
		buffer.WriteString(v.Sender)
		buffer.WriteString("&")
		buffer.WriteString(v.Receiver)
		buffer.WriteString("&")
		buffer.WriteString(v.TxAmount)
		buffer.WriteString("&")
		buffer.WriteString(v.PaymentMedium)
		buffer.WriteString("&")
		buffer.WriteString(v.TxNo)
		buffer.WriteString("&")
	}
	//TxArrayvalBytes,_ = json.Marshal(TxNewmap)
	//TxArraystr = buffer.String() 
	fmt.Println(TxNewmap)
	//fmt.Println(TxArraystr)

	//var TxValidArraystr string
	buffer.WriteString("&$$$$$$$&")
	for k,v := range TxValidNewmap{
		buffer.WriteString(k)
		buffer.WriteString("&")
		buffer.WriteString(v)
		buffer.WriteString("&")
	} 
	//TxValidArrayvalBytes,_ = json.Marshal(TxValidNewmap)
	TxArraystr = buffer.String() 
	fmt.Println(TxValidNewmap)
	//fmt.Println(TxValidArraystr)

	fmt.Println("Last ID :- %d",lastIDval)
	return shim.Success([]byte(TxArraystr))
}
/*
// Deletes an entity from state
func (t *SimpleChaincode) delete(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	A := args[0]

	// Delete the key from the state in ledger
	err := stub.DelState(A)
	if err != nil {
		return shim.Error("Failed to delete state")
	}

	return shim.Success(nil)
}

// query callback representing the query of a chaincode
func (t *SimpleChaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var A string // Entities
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the person to query")
	}

	A = args[0]

	// Get the state from the ledger
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + A + "\"}"
		return shim.Error(jsonResp)
	}

	if Avalbytes == nil {
		jsonResp := "{\"Error\":\"Nil amount for " + A + "\"}"
		return shim.Error(jsonResp)
	}

	jsonResp := "{\"Name\":\"" + A + "\",\"Amount\":\"" + string(Avalbytes) + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)
	return shim.Success(Avalbytes)
}  */

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
