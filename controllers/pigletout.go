package controllers

import (
	"github.com/astaxie/beego"
)

type PigletoutController struct {
	beego.Controller
}

func (c *PigletoutController) Get() {
	//登录检查

	//c.Data["IsShouy"] = true
	woshu := "123"
	c.Data["woshu"] = woshu

	c.TplName = "pigletout.html"

}
