package routers

import (
	"github.com/astaxie/beego"
	"myproject/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:GetData")
	beego.Router("get", &controllers.BlogController{}, "get:GetBlog")
	beego.Router("getAll", &controllers.BlogController{}, "get:GetAll")
	beego.Router("GetByIdAndDelete", &controllers.BlogController{}, "get:GetByIdAndDelete")
	beego.Router("delete/:id([0-9]+)", &controllers.BlogController{}, "get:DeleteBlog")

	beego.Router("/new", &controllers.MainController{}, "get:NewBlog")
	beego.Router("/add", &controllers.BlogController{}, "post:AddBlog")

	beego.Router("/update", &controllers.BlogController{}, "post:UpdateBlog")

	//beego.InsertFilter("/*", beego.BeforeRouter, &controllers.FilterUser)

	// //显示博客首页
	// beego.RegisterController("/", &controllers.IndexController{})
	// //查看博客详细信息
	// beego.RegisterController("/view/:id([0-9]+)", &controllers.ViewController{})
	// //新建博客博文
	// beego.RegisterController("/new", &controllers.NewController{})
	// //删除博文
	// beego.RegisterController("/delete/:id([0-9]+)", &controllers.DeleteController{})
	// //编辑博文
	// beego.RegisterController("/edit/:id([0-9]+)", &controllers.EditController{})
}
