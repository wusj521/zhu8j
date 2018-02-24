package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"zhu8j/controllers"
	"zhu8j/models"
	_ "zhu8j/routers"
)

func init() {
	// 注册数据库
	models.RegisterDB()
}

func main() {
	orm.Debug = true                      //开启 ORM 调试模式
	orm.RunSyncdb("default", false, true) //自动建立新DB表

	//注册路由
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/home", &controllers.HomeController{})
	beego.Router("/sow", &controllers.SowController{})
	beego.Router("/sowview", &controllers.SowviewController{})
	beego.Router("/pigletbirth", &controllers.PigletbirthController{})
	beego.Router("/pigletbirthview", &controllers.PigletbirthviewController{})
	beego.Router("/pigletout", &controllers.PigletoutController{})
	beego.Router("/pigletoutview", &controllers.PigletoutviewController{})
	beego.Router("/deadpig", &controllers.DeadpigController{})
	beego.Router("/deadpigview", &controllers.DeadpigviewController{})
	//在养情况
	beego.Router("/zaiyanginfo", &controllers.ZaiyanginfoController{})

	beego.Run()
}
