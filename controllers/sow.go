package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type SowController struct {
	beego.Controller
}

func (c *SowController) Get() {
	//登录检查
	woshu := "123"
	c.Data["woshu"] = woshu

	c.TplName = "sow.html"

}

func (c *SowController) Post() {
	//读取表单数据
	erhao := c.Input().Get("erhao")
	date := c.Input().Get("date")
	select2 := c.Input().Get("select2")

	fmt.Println(erhao)
	fmt.Println(date)
	fmt.Println(select2)

	c.Redirect("/sow", 302)
	return
}
