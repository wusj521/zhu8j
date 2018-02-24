package controllers

import (
	"github.com/astaxie/beego"
)

type PigletbirthController struct {
	beego.Controller
}

func (c *PigletbirthController) Get() {
	//登录检查

	//c.Data["IsShouy"] = true

	woshu := "123"
	c.Data["woshu"] = woshu

	c.TplName = "pigletbirth.html"

}
