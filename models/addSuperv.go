package models

import "github.com/hyperledger/fabric-sdk-go/pkg/client/channel"

//模块任务:
//将质量信息写入区块链

func (this * Application)addSupervItem(args []string)(string ,error){
	// 上传数据到区块链
	var tempArgs [][]byte
	for i:=1;i<len(args);i++{
		tempArgs = append(tempArgs,[]byte(args[i]))
	}
	request := channel.Request{ChaincodeID:this.fabricsetup.ChainCodeID,Fcn:args[0],Args:tempArgs}
	response,err := this.fabricsetup.client.Execute(request)
	return string(response.TransactionID),err
}