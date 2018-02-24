package controllers

import (
	"github.com/astaxie/beego"
)

type PigletoutviewController struct {
	beego.Controller
}

func (c *PigletoutviewController) Get() {
	//登录检查

	//c.Data["IsShouy"] = true
	woshu := "123"
	c.Data["woshu"] = woshu

	c.TplName = "pigletoutview.html"

}
