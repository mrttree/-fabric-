package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"strconv"
)

type Installer struct {
}

func(t *Installer)Init(stub shim.ChaincodeStubInterface)pb.Response{

	return shim.Success([]byte("success"))

}

func(t *Installer)Invoke(stub shim.ChaincodeStubInterface)pb.Response{
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

	case "numChecked":
		//'{"Args": ["numChecked"，"TransitInfo.Id"，"Num"]}'
		return numChecked(args,stub)

	case "qualityChecked":
		return qualityChecked()

	default:
		return shim.Error("Incorrect Function!!!")
	}
}

//物流到达后数量确认
func numChecked(args []string,stub shim.ChaincodeStubInterface)pb.Response{
	//'{"Args": ["numChecked"，"TransitInfo.Id"，"Num"]}'
	if len(args) != 2{
		return shim.Error("Incorrect arguments...Expecting a key")
	}
	//获取输入的数据，并将其处理为正确的格式
	Id := args[0]
	number := args[1]

	num, err := strconv.Atoi(number)
	if err != nil{
		return shim.Error("Atoi Error...")
	}


	//从账本中获取到物流订单的数据，将其进行处理
	TransitInfoRece, err := stub.GetState(Id)
	if err != nil{
		return  shim.Error("GetState Error...")
	}
	Transit := &TransitInfo{}
	err = json.Unmarshal(TransitInfoRece,&Transit)

	if Transit.Num != num{
		return shim.Error("订单数量不匹配，请重新核对获联系生产厂商！")
	}

	Transit.checked = true

	return shim.Success([]byte("numChecked Success!"))

}

func qualityChecked()pb.Response{
	return shim.Success([]byte("qualityChecked Success!"))
}


func main(){
	err := shim.Start(new(Installer))
	if err != nil{
		fmt.Println("Installer链码启动失败")
	}
}