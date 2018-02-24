package controllers

import (
	"github.com/astaxie/beego"
)

type PigletbirthviewController struct {
	beego.Controller
}

func (c *PigletbirthviewController) Get() {
	//登录检查

	//c.Data["IsShouy"] = true

	woshu := "123"
	c.Data["woshu"] = woshu

	c.TplName = "pigletbirthview.html"

}
