package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"strconv"
	"time"
)

type Transiter struct {

}

func (t *Transiter)Init(stub shim.ChaincodeStubInterface)pb.Response {
	return shim.Success([]byte("Transiter ChainCode Init success"))
}

func (t *Transiter)Invoke(stub shim.ChaincodeStubInterface)pb.Response{
	fn,args := stub.GetFunctionAndParameters()

	switch fn {
	case "queryTransit"://查询物流情况
		if len(args) != 1{
			return shim.Error("Incorrect arguments...Expecting a key")
		}
		//根据订单的Id查询
		return QueryTransit(args[0], stub)
	case "submitTransitInfo"://提交物流情况
		if len(args) != 10{
			return shim.Error("Incorrect arguments...Expecting a key")
		}
		return submitTransitInfo(args,stub)
	default:
		return shim.Error("Incorrect Function!!!")
	}

}

func submitTransitInfo(args []string, stub shim.ChaincodeStubInterface)pb.Response{
	var timeLayoutStr = "2006-01-02 15:04:05"


	key := args[0]
	DepartureTime := args[1]
	ArrivalTime := args[2]
	DeparturePlace := args[3]
	Destination := args[4]
	DepositTime := args[5]
	Transport := args[6]
	Logistics := args[7]
	fee := args[8]
	Num := args[9]
	Index := args[10]

	Id,err := strconv.Atoi(key)
	num, err := strconv.Atoi(Num)
	if err != nil{
		return  shim.Error("strconv.Atoi err ...")
	}
	departureTime, err := time.Parse(timeLayoutStr,DepartureTime)
	arrivalTime,err := time.Parse(timeLayoutStr,ArrivalTime)
	if err != nil{
		return  shim.Error("time.Parse err ...")
	}

	TransitInfo := &TransitInfo{
		Id:             Id,
		DepartureTime:  departureTime,
		ArrivalTime:    arrivalTime,
		DeparturePlace: DeparturePlace,
		Destination:    Destination,
		DepositTime:    DepositTime,
		Transport:      Transport,
		Logistics:      Logistics,
		fee:            fee,
		Num:            num,
		Index:          Index,
	}

	transiterInfo, err := json.Marshal(TransitInfo)
	if err != nil{
		return  shim.Error("json.Marshal err ...")
	}
	err = stub.PutState(key,transiterInfo)
	if err != nil{
		return  shim.Error("PutState err ...")
	}

	return shim.Success([]byte("submitTransitInfo Success ..."))
}




func main(){
	err := shim.Start(new(Transiter))
	if err != nil{
		fmt.Println("FirstParty链码启动失败")
	}
}