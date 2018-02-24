package controllers

import (
	"github.com/astaxie/beego"
)

type ZaiyanginfoController struct {
	beego.Controller
}

func (c *ZaiyanginfoController) Get() {
	//登录检查

	//c.Data["IsShouy"] = true
	woshu := "123"
	c.Data["woshu"] = woshu

	c.TplName = "zaiyanginfo.html"

}
