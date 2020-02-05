package main

import (
	"encoding/json"
	"fmt"
	pb "github.com/hyperledger/fabric/protos/peer"
	"strconv"
	"time"
)
import "github.com/hyperledger/fabric/core/chaincode/shim"

type PartProducer struct {

}

//母线相关零件生产商的智能合约
//查询母线/零件的情况
//查询物流情况
//提交零件的出厂情况  '{"Args": ["createPart"，part]}'
//生成实例化的零件实体，在结构体中插入相关数据，而后将其编码为json字符串，而后其他单位将其解码后再填充其他细节
//提交零件的物流情况
func (t *PartProducer)Init(stub shim.ChaincodeStubInterface)pb.Response{
	return  shim.Success([]byte("PartProducer ChainCode Init success"))
}

func (t *PartProducer)Invoke(stub shim.ChaincodeStubInterface)pb.Response{
	//获取参数
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

	case "submitPartInfo":
		if len(args) != 6{
			return shim.Error("Incorrect arguments...Expecting a key or value")
		}

		return submitPart(args,stub)

	default:
		return shim.Error("Incorrect Function!!!")
	}

}

func submitPart(args []string, stub shim.ChaincodeStubInterface)pb.Response{
	var timeLayoutStr = "2006-01-02 15:04:05"

	key := args[0]
	Name := args[1]
	Spec := args[2]
	MFRDate := args[3]
	MFRName := args[4]
	OutDate := args[5]

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

	Component := &Component{
		Id:             Id,
		Name:           Name,
		Spec:           Spec,
		MFGDate:        mfrDate,
		MFRSName:       MFRName,
		OutDate:        outDate,
	}

	component, err := json.Marshal(Component)
	if err != nil{
		return  shim.Error("json.Marshal err ...")
	}
	err = stub.PutState(key,component)
	if err != nil{
		return  shim.Error("PutState err ...")
	}

	return shim.Success([]byte("submitGenerator success!"))

}

func main(){
	err := shim.Start(new(PartProducer))
	if err != nil{
		fmt.Println("FirstParty链码启动失败")
	}
}