package controllers

import (
	//"zhubajie/models"

	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Get() {
	//登录检查

	//c.Data["IsShouy"] = true

	woshu := "123"
	c.Data["woshu"] = woshu

	c.TplName = "index.html"

	/*	c.Data["Website"] = "beego.me"
		c.Data["Email"] = "astaxie@gmail.com"
		c.TplName = "index.tpl"

		c.Data["TrueCond"] = true
		c.Data["FalseCond"] = false

		type u struct {
			Name string
			Age  int
			Sex  string
		}
		User := &u{
			Name: "wusj",
			Age:  36,
			Sex:  "Male",
		}
		c.Data["User"] = User

		numb := []int{1, 2, 3, 4, 5, 6}
		c.Data["Numb"] = numb
	*/

}
