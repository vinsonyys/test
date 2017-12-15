package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) GetData() {
	c.Data["Website"] = "beego.me" //存储输出数据的map

	c.Data["Email"] = "astaxie@gmail.com"

	c.TplName = "index.tpl" //  喧染的模板
}

func (c *MainController) NewBlog() {
	c.TplName = "new.tpl"
}
