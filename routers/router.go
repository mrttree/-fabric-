package routers

import (
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    // 添加信息路由

    // 添加物流信息
	beego.Router("/form_transit",&controllers.MainController{},"get:FormTransit")
    //添加零件信息
	beego.Router("/form_part",&controllers.MainController{},"get:FormPart")
    //添加母线信息
	beego.Router("/form_produce",&controllers.MainController{},"get:FormPooduce")
    //添加质监信息
	beego.Router("/form_superv",&controllers.MainController{},"get:FormSuperv")


    // 查询信息路由

    //查询物流信息
	beego.Router("/transit_search",&controllers.MainController{},"get:TransitSearch")
    //查询母线信息
	beego.Router("/produce_search",&controllers.MainController{},"get:ProduceSearch")
    //查询零件信息
	beego.Router("/part_search",&controllers.MainController{},"get:PartSearch")

}
