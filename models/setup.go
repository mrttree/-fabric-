package models

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/fab/ccpackager/gopackager"
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/common/cauthdsl"
	"github.com/pkg/errors"
)

// pkg/client/channel : 提供与channel交易相关的一些操作
// pkg/client/event   : 提供一些访问 Fabric 网络上的通道时间,比如 交易时间,区块的产生...
// pkg/client/resmgmt : 提供资源管理功能,例如 链代码的安装,channel的创建
// pkg/client/msp	  : 提供与身份管理相关的功能


// 这个文件中得到 Fabric_sdk
func (this *FabricSetup)Initialize() error{
	// 1.根据配置文件 得到SDK 实例
	sdk,err := fabsdk.New(config.FromFile(this.ConfigFile))
	if err!=nil{
		return errors.WithMessage(err,"Failed to create SDK")
	}
	this.sdk = sdk
	fmt.Println("SDK created!")

	// 2. 通过 SDK 实例,基于用户和组织创建上下文[为 创建 resMgmtCli 做准备]
	resCliMgmtcontext:= this.sdk.Context(fabsdk.WithUser(this.OrgAdmin),fabsdk.WithOrg(this.OrgName))

	// 3. 创建资源管理客户端 - resMgmtCli
	resMgmtCli ,err := resmgmt.New(resCliMgmtcontext)
	if err!=nil{
		return errors.WithMessage(err,"Create resMgmtCli error!...")
	}
	this.resMgmtCli = resMgmtCli

	fmt.Println("Create resource managment client success!")

	// 4. 利用资源管理客户端,创建 channel,得到一个操作链代码的 client
	request := resmgmt.SaveChannelRequest{
		ChannelID:this.channelID,
		ChannelConfigPath:this.channelConfig,
	}
	txID,err := this.resMgmtCli.SaveChannel(request)
	if err!=nil||txID.TransactionID == ""{
		return errors.WithMessage(err,"SaveChannel error!")
	}
	fmt.Println("channel created!")

	// 把组织添加到 channel 中的时候,一般制定一些重试的策略,和指定 orderer节点的网络位置
	err = this.resMgmtCli.JoinChannel(
		this.channelID,
		resmgmt.WithRetry(retry.DefaultResMgmtOpts),
		// orderer.itcast.cn
		resmgmt.WithOrdererEndpoint(this.OrdererID),
		)

	if err!=nil{
		return errors.WithMessage(err,"Failed to join channel!")
	}
	fmt.Println("Joined channel!")

	fmt.Println("Initialize Successful!")

	return nil
}

// 安装链代码
func (this *FabricSetup)InstallAndInstantiate()error{

	// 准备 安装链代码所需要的参数
	ccPkg,err := gopackager.NewCCPackage(
		this.ChaincodePath,
		this.ChaincodeGoPath,
		)
	if err!=nil{
		fmt.Println("Chaincode packager error!")
		return err
	}

	installResqust := resmgmt.InstallCCRequest{
		Name:this.ChainCodeID,
		Path:"chaincode",
		Version:"1.0",
		Package:ccPkg,
	}

	// 1. 安装链代码
	_,err = this.resMgmtCli.InstallCC(installResqust)
	if err!=nil{
		fmt.Println("chaincode install error!")
	}

	fmt.Println("chaincode install success!")

	// 2. 实例化链代码,在实例化之前,我们需要实现指定一个背书策略
	// 组织的MSP
	ccpolity := cauthdsl.SignedByAnyMember([]string{"ofgj.itcast.cn"})

	txID,err := this.resMgmtCli.InstantiateCC(
		this.channelID,
		resmgmt.InstantiateCCRequest{
			Name:this.ChainCodeID,
			Path:this.ChaincodeGoPath,
			Version:"1.0",
			Args:nil,
			Policy:ccpolity,
		},
	)
	if err!=nil|| txID.TransactionID ==""{
		fmt.Println("Instantiate chaincode error!")
		return err
	}

	fmt.Println("chaincode instantiate success!")


	// 实例化完成之后,就可以操作账本数据了,我们还需要,在创建一个 channel 的客户端

	chCliContext := this.sdk.ChannelContext(this.channelID,fabsdk.WithUser(this.UserName))

	client,err := channel.New(chCliContext)
	this.client = client

	fmt.Println("channel client created!")

	fmt.Println("chaincode install & inxtantiate successful!")

	return nil
}

// 5. 资源的释放
func (this *FabricSetup) CloseSDK(){
	this.sdk.Close()
}
