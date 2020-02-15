package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"strconv"
	"time"
)

type Installer struct {
}

type Generator struct {                //母线定义
	Id int                              //编号
	Name string                         //名称
	Spec string                         //规格
	MFGDate time.Time                   //生产日期
	MFRSName string                     //生产厂商
	OutDate time.Time                   //出厂日期
	BatchNum int                        //批次号及合同编号
	ComponentId string                  //零件,生产的零件编号组合键
	//上面部分由母线生产厂商写入，单独背书
	TransitInfo  TransitInfo            //运输信息
	//由母线厂家及运输单位写入，共同背书
	QualityCertifi QualityCertifi       //质量证明证书
	//复杂的背书策略
	Checked    Examine                  //验收验证
	MaintenInfo MaintenInfo             //维护信息
}

type Component struct {                 //零件信息
	Id int                              //编号
	Name string                         //名称
	Spec string                         //规格
	MFGDate time.Time               //生产日期
	MFRSName string                     //生产厂商
	OutDate time.Time               //出厂日期
	BatchNum int                        //批次号及合同编号
	//运输信息
	QualityCertifi QualityCertifi       //质量证明证书
	Checked    Examine                  //质量验证
}

type QualityCertifi struct {//质量验证
	//编号
	//名称
	//质量认证证书
	//产品试验报告
	//母线生产厂家质量体系认证证书
	//质量验证单位（多家单位）
}

type Examine struct {//检查记录

	//到货质量初验收
	//安装后质量验收

}

type MaintenInfo struct {//维护信息
	//日常使用维护记录
	//售后记录
}

type TransitInfo struct {
	//母线及零件的物流信息依据不同的链来隔离
	Id int                     //运输订单编号
	DepartureTime time.Time    //出发时间
	ArrivalTime time.Time      //到达时间
	DeparturePlace string      //出发地
	Destination string         //目的地
	DepositTime string         //中转存储时间
	Transport string           //运送方式
	Logistics string           //物流公司信息
	fee string                 //费用
	Num int                    //母线或零件数量
	checked bool               //到货后数量验收
	Index string                  //母线id的集合

}

func(t *Installer)Init(stub shim.ChaincodeStubInterface)pb.Response{

	fmt.Println("============install chaincode init success=============")

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

	case "submitGeneratorInfo":
		if len(args) != 7{
			return shim.Error("Incorrect arguments...Expecting a key or value")
		}

		//根据key value 插入母线信息
		return submitGenerator(args,stub)

	case "submitPartInfo":
		if len(args) != 6{
			return shim.Error("Incorrect arguments...Expecting a key or value")
		}

		return submitPart(args,stub)

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



//查询过程基本相同，后期可封装

func QueryGenerator(args string,stub shim.ChaincodeStubInterface)pb.Response{

	//generator,err := stub.GetState(args)
	//return shim.Success(generator)
	//getstate安全性不如gethistoryforkey
	//if err != nil{
	//	return shim.Error("GetState Error...")
	//}
	if len(args) != 1{
		return shim.Error("Incorrect number of arguments.")
	}

	ID := args[0]

	//获取查询结果迭代器
	resultItreator, err := stub.GetHistoryForKey(ID)
	if err != nil{
		shim.Error("gethistoryforKey err...")
	}
	defer resultItreator.Close()

	var GeneratorInfo Generator

	//使用迭代器遍历查询结果集
	for resultItreator.HasNext(){
		var generatorInfo Generator

		response, err := resultItreator.Next()
		if err != nil {
			return shim.Error("range.Next err...")
		}

		err = json.Unmarshal(response.Value,&generatorInfo)
		if err!=nil{
			return shim.Error("json.unmarshal err...")
		}

		GeneratorInfo = generatorInfo
	}

	jsonAsBytes, err  := json.Marshal(GeneratorInfo)
	if err!=nil{
		return shim.Error("json.marshal err...")
	}

	return shim.Success(jsonAsBytes)


}

func QueryPart(args string,stub shim.ChaincodeStubInterface)pb.Response{
	//generator,err := stub.GetState(args)
	//return shim.Success(generator)
	//getstate安全性不如gethistoryforkey
	//if err != nil{
	//	return shim.Error("GetState Error...")
	//}
	if len(args) != 1{
		return shim.Error("Incorrect number of arguments.")
	}

	ID := args[0]

	//获取查询结果迭代器
	resultItreator, err := stub.GetHistoryForKey(ID)
	if err != nil{
		shim.Error("gethistoryforKey err...")
	}
	defer resultItreator.Close()

	var Component Generator

	//使用迭代器遍历查询结果集
	for resultItreator.HasNext(){
		var component Generator

		response, err := resultItreator.Next()
		if err != nil {
			return shim.Error("range.Next err...")
		}

		err = json.Unmarshal(response.Value,&component)
		if err!=nil{
			return shim.Error("json.unmarshal err...")
		}

		Component = component
	}

	jsonAsBytes, err  := json.Marshal(Component)
	if err!=nil{
		return shim.Error("json.marshal err...")
	}

	return shim.Success(jsonAsBytes)
}

func QueryTransit(args string,stub shim.ChaincodeStubInterface)pb.Response {
	//generator,err := stub.GetState(args)
	//return shim.Success(generator)
	//getstate安全性不如gethistoryforkey
	//if err != nil{
	//	return shim.Error("GetState Error...")
	//}
	if len(args) != 1{
		return shim.Error("Incorrect number of arguments.")
	}

	ID := args[0]

	//获取查询结果迭代器
	resultItreator, err := stub.GetHistoryForKey(ID)
	if err != nil{
		shim.Error("gethistoryforKey err...")
	}
	defer resultItreator.Close()

	var TransitInfo Generator

	//使用迭代器遍历查询结果集
	for resultItreator.HasNext(){
		var transitInfo Generator

		response, err := resultItreator.Next()
		if err != nil {
			return shim.Error("range.Next err...")
		}

		err = json.Unmarshal(response.Value,&transitInfo)
		if err!=nil{
			return shim.Error("json.unmarshal err...")
		}

		TransitInfo = transitInfo
	}

	jsonAsBytes, err  := json.Marshal(TransitInfo)
	if err!=nil{
		return shim.Error("json.marshal err...")
	}

	return shim.Success(jsonAsBytes)
}


func main(){
	err := shim.Start(new(Installer))
	if err != nil{
		fmt.Println("Installer链码启动失败")
	}
}
