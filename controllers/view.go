package controllers

import (
	"github.com/astaxie/beego"
	//"myproject/models"
	//"strconv"
)

type ViewController struct {
	beego.Controller
}

// func (this *ViewController) Get() {
// 	id, _ := strconv.Atoi(this.Ctx.Input.Params(":id"))
// 	this.Data["Post"] = models.GetBlog(id)
// 	this.Layout = "layout.tpl"
// 	this.TplNames = "view.tpl"
// 	this.Render()
// }
