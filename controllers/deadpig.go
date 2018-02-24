package controllers

import (
	"github.com/astaxie/beego"
)

type DeadpigController struct {
	beego.Controller
}

func (c *DeadpigController) Get() {
	//登录检查

	//c.Data["IsShouy"] = true

	woshu := "123"
	c.Data["woshu"] = woshu

	c.TplName = "deadpig.html"

}
