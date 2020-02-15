package models

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"time"
)

// addHouse
// addOrder

type Application struct{
	// 从这个属性身上可以拿到 操作 账本数据的句柄
	fabricsetup *FabricSetup
}

type FabricSetup struct {
	ConfigFile  string
	ChainCodeID string
	ChaincodePath string
	ChaincodeGoPath string
	OrgAdmin string
	OrgName  string
	UserName string
	channelID string
	channelConfig string
	OrdererID string
	sdk        *fabsdk.FabricSDK
	resMgmtCli *resmgmt.Client
	client     *channel.Client
}

// 维护 软件的数据模型

type Generator struct {                //母线定义
	Id int                              //编号
	Name string                         //名称
	Spec string                         //规格
	MFGDate time.Time               //生产日期
	MFRSName string                     //生产厂商
	OutDate time.Time               //出厂日期
	BatchNum int                        //批次号及合同编号
	ComponentId string                   //零件,生产的零件编号组合键
	//上面部分由母线生产厂商写入，单独背书
	TransitInfo  TransitInfo            //运输信息
	//由母线厂家及运输单位写入，共同背书
	QualityCertifi QualityCertifi       //质量证明证书
	//复杂的背书策略
	Checked    Examine                  //质量验证
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
	Id int//质量认证编号
	Name string//名称
	certificate string//质量认证证书,后期可换为证书证明二进制文件的哈希
	reprort string//产品试验报告
	ProducerCertificate string//母线生产厂家质量体系认证证书
	sign string//质量验证单位（多家单位）
}

type Examine struct {//检查记录

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
	producerChecked bool       //厂家确认物流订单
	installerChecked bool      //到货后数量验收
	Index string                  //母线id的集合

}