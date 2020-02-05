package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"strconv"
	"time"
)

type Producer struct {

}

//母线生产商的智能合约
//查询母线/零件的情况
//查询物流情况
//提交母线的出厂情况  '{"Args": ["createGenerator"，Generator]}'
//生成实例化的母线实体，在结构体中插入相关数据，而后将其编码为json字符串，而后其他单位将其解码后再填充其他细节
//提交母线的物流情况
func (t *Producer)Init(stub shim.ChaincodeStubInterface)pb.Response{
	return  shim.Success([]byte("Producer ChainCode Init success"))
}

func (t *Producer)Invoke(stub shim.ChaincodeStubInterface)pb.Response{
	//获取参数
	fn, args := stub.GetFunctionAndParameters()
	//校验参数
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

	case "submitGeneratorInfo":
		if len(args) != 7{
			return shim.Error("Incorrect arguments...Expecting a key or value")
		}

		//根据key value 插入母线信息
		return submitGenerator(args,stub)

	default:
		return shim.Error("Incorrect Function!!!")
	}

}

func submitGenerator(args []string, stub shim.ChaincodeStubInterface)pb.Response {
	var timeLayoutStr = "2006-01-02 15:04:05"

	key := args[0]
	Name := args[1]
	Spec := args[2]
	MFRDate := args[3]
	MFRName := args[4]
	OutDate := args[5]
	ComponentId := args[6]

	Id,err := strconv.Atoi(key)
	if err != nil{
		return  shim.Error("strconv.Atoi err ...")
	}

	mfrDate, err := time.Parse(timeLayoutStr,MFRDate)
	if err != nil{
		return  shim.Error("time.Parse err ...")
	}
	outDate, err := time.Parse(timeLayoutStr,OutDate)
	if err != nil{
		return  shim.Error("time.Parse err ...")
	}
	Generator := &generator{
	 	Id: Id,
	 	Name: Name,
	 	Spec: Spec,
	 	MFGDate: mfrDate,
	 	MFRSName:MFRName,
	 	OutDate:outDate,
	 	ComponentId:ComponentId,
	}

	generator,err := json.Marshal(Generator)
	if err != nil{
		return  shim.Error("json.Marshal err ...")
	}
	err =stub.PutState(key,generator)
	if err != nil{
		return  shim.Error("PutState err ...")
	}

	return shim.Success([]byte("submitGenerator success!"))


}

func main(){
	err := shim.Start(new(Producer))
	if err != nil{
		fmt.Println("FirstParty链码启动失败")
	}
}