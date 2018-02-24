package controllers

import (
	"github.com/astaxie/beego"
)

type DeadpigviewController struct {
	beego.Controller
}

func (c *DeadpigviewController) Get() {
	//登录检查

	//c.Data["IsShouy"] = true

	woshu := "123"
	c.Data["woshu"] = woshu

	c.TplName = "deadpigview.html"

}
