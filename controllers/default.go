package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"fmt"
	"Generator-Anti-counterfeiting-project/models"
)

type MainController struct {
	beego.Controller
}


func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
}
//物流及母线接受单位（安装单位）填写
func (this *MainController)FormTransit(){
	//相应请求，页面返回客户端
	this.TplName = "form_transit.html"

	transitID := this.GetString("transitID")
	if transitID == "" {
		return
	} else {
		// 获取用户输入
		starttime := this.GetString("starttime")
		arrivaltime := this.GetString("arrivaltime")
		departurePlace := this.GetString("departurePlace")
		destination := this.GetString("destination")
		depositTime := this.GetString("depositTime")
		transport := this.GetString("transport")
		logistics := this.GetString("logistics")
		fee := this.GetString("fee")
		num := this.GetString("num")
		index := this.GetString("index")

		// 组织参数
		var args []string
		args = append(args, "submitTransitInfo")
		args = append(args, starttime)
		args = append(args, arrivaltime)
		args = append(args, departurePlace)
		args = append(args, destination)
		args = append(args, depositTime)
		args = append(args, transport)
		args = append(args, logistics)
		args = append(args, fee)
		args = append(args, num)
		args = append(args, index)


		ret, err := models.App.AddTransitItem(args)
		if err != nil {
			fmt.Println("AddTransitItem err...")
		}
		fmt.Println("<--- 添加物流信息结果　--->：", ret)
	}

	this.TplName = "index.html"

}
//由零件供应商添加
func (this *MainController)FormPart(){
	//相应请求，页面返回客户端
	this.TplName = "form_transit.html"

	partID := this.GetString("partID")
	if partID == "" {
		return
	} else {
		// 获取用户输入
		name := this.GetString("name")
		spec := this.GetString("spec")
		mFRDate := this.GetString("mFRDate")
		mFRName := this.GetString("mFRName")
		outDate := this.GetString("outDate")



		// 组织参数
		var args []string
		args = append(args, "submitGeneratorInfo")
		args = append(args, partID)
		args = append(args, name)
		args = append(args, spec)
		args = append(args, mFRDate)
		args = append(args, mFRName)
		args = append(args, outDate)



		ret, err := models.App.AddTransitItem(args)
		if err != nil {
			fmt.Println("AddTransitItem err...")
		}
		fmt.Println("<--- 添加母线信息结果　--->：", ret)
	}

	this.TplName = "index.html"
}
//由母线供应商添加
func (this *MainController)FormPooduce(){
	//相应请求，页面返回客户端
	this.TplName = "form_transit.html"

	produceID := this.GetString("produceID")
	if produceID == "" {
		return
	} else {
		// 获取用户输入
		name := this.GetString("name")
		spec := this.GetString("spec")
		mFRDate := this.GetString("mFRDate")
		mFRName := this.GetString("mFRName")
		outDate := this.GetString("outDate")
		componentId := this.GetString("componentId")


		// 组织参数
		var args []string
		args = append(args, "submitGeneratorInfo")
		args = append(args, produceID)
		args = append(args, name)
		args = append(args, spec)
		args = append(args, mFRDate)
		args = append(args, mFRName)
		args = append(args, outDate)
		args = append(args, componentId)



		ret, err := models.App.AddTransitItem(args)
		if err != nil {
			fmt.Println("AddTransitItem err...")
		}
		fmt.Println("<--- 添加母线信息结果　--->：", ret)
	}

	this.TplName = "index.html"
}
//由质量监察监督单位添加
func (this *MainController)FormSuperv(){

}
//物流信息查询
func (this *MainController)TransitSearch(){
	//获取用户需要查询的transitID
    this.TplName = "transit_search.html"
    //组合链码需要的参数
    key := this.GetString("transitID")
    var args []string
    args = append(args,"queryTransit")
    args = append(args,key)

    //调用model层的函数，查询数据
    response, err := models.App.GetTransitINfo(args)
    if err != nil{
    	fmt.Println("models.App.GetTransitINfo err..")
	}

	var jsonData models.TransitInfo
    err = json.Unmarshal([]byte(response),&jsonData)
    if err != nil{
    	fmt.Println("json.Unmarshal")
	}

	fmt.Println("========jsonData==========",jsonData)

	this.Data["transitId"] = jsonData.Id
	this.Data["departureTime"] = jsonData.DepartureTime
	this.Data["arrivalTime"] = jsonData.ArrivalTime
	this.Data["departurePlace"] = jsonData.DeparturePlace
	this.Data["destination"] = jsonData.Destination
	this.Data["depositTime"] = jsonData.DepositTime
	this.Data["transport"] = jsonData.Transport
	this.Data["logistics"] = jsonData.Logistics
	this.Data["num"] = jsonData.Num
}
//母线信息查询
func (this *MainController)ProduceSearch(){
	//获取用户需要查询的母线ID
	this.TplName = "produce_search.html"
	//组合链码需要的参数
	key := this.GetString("produceID")
	var args []string
	args = append(args,"queryProduce")
	args = append(args,key)

	//调用model层的函数，查询数据
	response, err := models.App.GetProduceINfo(args)
	if err != nil{
		fmt.Println("models.App.GetProduceINfo err..")
	}

	var jsonData models.Generator
	err = json.Unmarshal([]byte(response),&jsonData)
	if err != nil{
		fmt.Println("json.Unmarshal")
	}

	fmt.Println("========jsonData==========",jsonData)

	this.Data["produceID"] = jsonData.Id
	this.Data["Name"] = jsonData.Name
	this.Data["Spec"] = jsonData.Spec
	this.Data["MFGDate"] = jsonData.MFGDate
	this.Data["MFRSName"] = jsonData.MFRSName
	this.Data["OutDate"] = jsonData.OutDate
	this.Data["BatchNum"] = jsonData.BatchNum
	this.Data["ComponentId"] = jsonData.ComponentId
	this.Data["TransitInfo"] = jsonData.TransitInfo
	this.Data["QualityCertifi"] = jsonData.QualityCertifi
}
//零件信息查询
func (this *MainController)PartSearch(){
	//获取用户需要查询的零件ID
	this.TplName = "part_search.html"
	//组合链码需要的参数
	key := this.GetString("partID")
	var args []string
	args = append(args,"queryPart")
	args = append(args,key)

	//调用model层的函数，查询数据
	response, err := models.App.GetPartINfo(args)
	if err != nil{
		fmt.Println("models.App.GetProduceINfo err..")
	}

	var jsonData models.Component
	err = json.Unmarshal([]byte(response),&jsonData)
	if err != nil{
		fmt.Println("json.Unmarshal")
	}

	fmt.Println("========jsonData==========",jsonData)

	this.Data["partID"] = jsonData.Id
	this.Data["Name"] = jsonData.Name
	this.Data["Spec"] = jsonData.Spec
	this.Data["MFGDate"] = jsonData.MFGDate
	this.Data["MFRSName"] = jsonData.MFRSName
	this.Data["OutDate"] = jsonData.OutDate
	this.Data["BatchNum"] = jsonData.BatchNum
	this.Data["QualityCertifi"] = jsonData.QualityCertifi
}