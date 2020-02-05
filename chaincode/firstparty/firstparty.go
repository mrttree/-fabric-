package main

import (
	"fmt"
	pb "github.com/hyperledger/fabric/protos/peer"
	//"time"
)
import "github.com/hyperledger/fabric/core/chaincode/shim"

// 甲方的智能合约，查询零件/母线的情况  '{"Args": ["queryGenerator"，"generatorId"]}' '{"Args": ["queryPart"，"PartId"]}'
//查询物流情况                       '{"Args": ["queryTransit"，"OrderId"]}'
//（查询质量情况）
//（提交售后情况）
type FirstParty struct {
}

func(t *FirstParty)Init(stub shim.ChaincodeStubInterface)pb.Response{

	return shim.Success([]byte("success"))

}

func(t *FirstParty)Invoke(stub shim.ChaincodeStubInterface)pb.Response{
	//接受参数
	fn, args := stub.GetFunctionAndParameters()

	switch fn {
	case "queryGenerator":
		if len(args) != 1{
			return shim.Error("Incorrect arguments...Expecting a key")
		}
		//根据母线的Id查询
		return QueryGenerator(args[0], stub)

	case "queryPart":
		if len(args) != 1{
			return shim.Error("Incorrect arguments...Expecting a key")
		}
		//根据零件的Id查询
		return QueryPart(args[0], stub)

	case "queryTransit"://查询物流情况
		if len(args) != 1{
			return shim.Error("Incorrect arguments...Expecting a key")
		}
		//根据订单的Id查询
		return QueryTransit(args[0], stub)


	default:
		return shim.Error("Incorrect Function!!!")
	}

}


func main(){
	err := shim.Start(new(FirstParty))
	if err != nil{
		fmt.Println("FirstParty链码启动失败")
	}
}