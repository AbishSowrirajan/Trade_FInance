/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

/*
 * The sample smart contract for documentation topic:
 * Trade Finance Use Case - WORK IN  PROGRESS
 */

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

// Define the letter of credit
type LetterOfCredit struct {
	LCId       string `json:"lcId"`
	ExpiryDate string `json:"expiryDate"`
	Buyer      string `json:"buyer"`
	ImBank     string `json:"imbank"`
	Seller     string `json:"seller"`
	ExBank     string `jsin:"exbank"`
	Amount     int    `json:"amount"`
	Status     string `json:"status"`
}

// Define the shipment status 
type Shipment struct {
	ShipmentID string `json:"ShipmentID"`
	LCId       string `json:"lcId"`
	Description string `json:"Description"`
	ShipmentValue      string `json:"ShipmentValue "`
	Exbank       string `json:"Exbank"`
	Imbank     string `json:"Imbank"`
	ShipmentCo    string    `json:"ShiptmentCo"`
	Poland     string `json:"Poland"`
	Poentry    string `json:"Poentry"`
	Shipmentstatus string `json:"Shipmentstatus"`
}
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "requestLC" {
		return s.requestLC(APIstub, args)
	} else if function == "issueLC" {
		return s.issueLC(APIstub, args)
	} else if function == "acceptLC" {
		return s.acceptLC(APIstub, args)
	} else if function == "exporter" {
		return s.exporter(APIstub, args)
	} else if function == "getLC" {
		return s.getLC(APIstub, args)
	} else if function == "getLCHistory" {
		return s.getLCHistory(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

// This function is initiate by Buyer
func (s *SmartContract) requestLC(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	lcId := args[0]
	expiryDate := args[1]
	buyer := args[2]
	imbank := args[3]
	seller := args[4]
	exbank := arg[5]
	amount, err := strconv.Atoi(args[6])
	if err != nil {
		return shim.Error("Not able to parse Amount")
	}

	LC := LetterOfCredit{LCId: lcId, ExpiryDate: expiryDate, Buyer: buyer, ImBank: imbank, Seller: seller,ExBank:exbank Amount: amount, Status: "Requested"}
	LCBytes, err := json.Marshal(LC)

	APIstub.PutState(lcId, LCBytes)
	fmt.Println("LC Requested -> ", LC)

	return shim.Success(nil)
}

// This function is initiate by Seller
func (s *SmartContract) issueLC(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	lcId := args[0]

	// if err != nil {
	// 	return shim.Error("No Amount")
	// }

	LCAsBytes, _ := APIstub.GetState(lcId)

	var lc LetterOfCredit

	err := json.Unmarshal(LCAsBytes, &lc)

	if err != nil {
		return shim.Error("Issue with LC json unmarshaling")
	}

	LC := LetterOfCredit{LCId: lc.LCId, ExpiryDate: lc.ExpiryDate, Buyer: lc.Buyer, ImBank: lc.ImBank, Seller: lc.Seller,ExBank:lc.ExBank, Amount: lc.Amount, Status: "Issued"}
	LCBytes, err := json.Marshal(LC)

	if err != nil {
		return shim.Error("Issue with LC json marshaling")
	}

	APIstub.PutState(lc.LCId, LCBytes)
	fmt.Println("LC Issued -> ", LC)

	return shim.Success(nil)
}

func (s *SmartContract) acceptLC(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	lcId := args[0]

	LCAsBytes, _ := APIstub.GetState(lcId)

	var lc LetterOfCredit

	err := json.Unmarshal(LCAsBytes, &lc)

	if err != nil {
		return shim.Error("Issue with LC json unmarshaling")
	}

	LC := LetterOfCredit{LCId: lc.LCId, ExpiryDate: lc.ExpiryDate, Buyer: lc.Buyer, ImBank: lc.ImBank, Seller: lc.Seller,ExBank:lc.ExBank, Amount: lc.Amount, Status: "Accepted"}
	LCBytes, err := json.Marshal(LC)

	if err != nil {
		return shim.Error("Issue with LC json marshaling")
	}

	APIstub.PutState(lc.LCId, LCBytes)
	fmt.Println("LC Accepted -> ", LC)

	return shim.Success(nil)
}

func (s *SmartContract) exporter(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	lcId := args[0]
	shipmentId := args[1]
	description := args[2]
	shipmentValue := args[3]
	//exbank := args[4]
	//imbank :=args[5]
	shipmentCo := arg[4]
	poland := arg[5]
	poentry := arg[6]
	shipmentstaus := "informed"



	LCAsBytes, _ := APIstub.GetState(lcId)

	var lc LetterOfCredit

	err := json.Unmarshal(LCAsBytes, &lc)

	if err != nil {
		return shim.Error("Issue with LC json unmarshaling")
	}

	if lc.LCId == lcId & lc.Status == "accepted" {

		retun shim.error("Letter of credit is not accepted")
	}

	SS := Shipment{ShipmentId:shipmentId,LcId:lcId,Description:description,ShipmentValue:shipmentValue,
	               Exbank:lc.ExBank,Imbank:lc.ImBank,ShipmentCo:shipmentco,Poland:poland,Poentry:poentry,Shipmentstatus:shipmentstaus}
	SSBytes, err := json.Marshal(SS)
	
	if err != nil {
		return shim.Error("Issue with Shipment json marshaling")
	}
	
	APIstub.PutState(shipmentId, SSBytes)
	fmt.Println("Shipment process Initiated by Exporter -> ")

	return shim.Success(nil)
}

func (s *SmartContract) getLC(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	lcId := args[0]

	// if err != nil {
	// 	return shim.Error("No Amount")
	// }

	LCAsBytes, _ := APIstub.GetState(lcId)

	return shim.Success(LCAsBytes)
}

func (s *SmartContract) getLCHistory(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	lcId := args[0]

	resultsIterator, err := APIstub.GetHistoryForKey(lcId)
	if err != nil {
		return shim.Error("Error retrieving LC history.")
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing historic values for the marble
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error("Error retrieving LC history.")
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

	fmt.Printf("- getLCHistory returning:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
