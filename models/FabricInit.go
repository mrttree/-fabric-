package models

import (
	"fmt"
	"os"
)

var App Application

// init 一旦加载,所有与 创建 SDK 相关的一些配置 就被初始化
func init(){
	fSetup := FabricSetup{
		ConfigFile:"conf/config.yaml",
		OrgAdmin:"Admin",
		UserName:"User1",
		OrgName:"OrgInstall",
		ChainCodeID : "mycc_install",
		ChaincodePath:"origins/chaincode",
		ChaincodeGoPath:os.Getenv("GOPATH"),

		channelID:"SupervOrgchannel",
		channelConfig:"/home/ts/workspace/src/origin/conf/channel-artifacts/SupervOrgchannel.tx",

		OrdererID:"orderer.trace.com",
	}

	// 得到 SDK
	err := fSetup.Initialize()
	if err!=nil{
		fmt.Printf("Unable Initizlize SDK,%v!\n",err)
		return
	}

	// 安装链代码的操作
	err = fSetup.InstallAndInstantiate()
	if err!=nil{
		fmt.Printf("Unable InstallAndInstantiate,%v!\n",err)
		return
	}

	defer fSetup.CloseSDK()

	App = Application{
		fabricsetup:fSetup,
	}

}
