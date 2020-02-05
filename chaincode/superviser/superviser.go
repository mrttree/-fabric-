package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)


//质量监管单位，如政府质监局，由甲方聘请的第三方质量管理及监理公司
//在生产及安装过程中对产品质量进行检查
type superviser struct {

}

func (t *superviser)Init(stub shim.ChaincodeStubInterface)pb.Response{
	return  shim.Success([]byte("superviser ChainCode Init success"))
}

func (t *superviser)Invoke(stub shim.ChaincodeStubInterface)pb.Response{
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

	case "submitQualityChecked":
		return submitQualityChecked(args,stub)

	default:
		return shim.Error("Incorrect Function!!!")
	}


}

func submitQualityChecked(args []string,stub shim.ChaincodeStubInterface)pb.Response{

	return shim.Success([]byte("Submit Success ..."))
}




func main(){
	err := shim.Start(new(superviser))
	if err != nil{
		fmt.Println("Superviser链码启动失败")
	}
}
