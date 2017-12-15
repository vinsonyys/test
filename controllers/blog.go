package controllers

import (
	//"github.com/astaxie/beego"
	"fmt"
	"github.com/astaxie/beego/orm"
	"myproject/models"
	"strconv"
	"time"
)

type BlogController struct {
	MainController
}

func (this *BlogController) GetBlog() {
	//id, _ := this.GetInt("id")
	// var idaa int
	// this.Ctx.Input.Bind(&idaa, "id")
	// fmt.Println("--------", idaa)
	//id := this.Input().Get("id")
	fmt.Println("-------------" + this.Ctx.Input.Param("id"))
	id := this.GetString("id")
	intid, _ := strconv.Atoi(id)
	fmt.Println("=========", intid)
	success := false
	response := "0"
	var blog models.Blog

	if intid > 0 {
		o := orm.NewOrm()
		qs := o.QueryTable("Blog")
		err := qs.Filter("id", intid).One(&blog)
		if err == nil {
			success = true

		} else {
			response = "4001"
		}

	} else {
		response = "4000"
	}

	if success == true {
		this.Data["json"] = map[string]interface{}{"status": 0, "desc": response, "data": blog}
	} else {
		this.Data["json"] = map[string]interface{}{"status": 1, "desc": response}
	}
	this.ServeJSON()

}

func (this *BlogController) GetAll() {
	success := false
	respone := "0"
	var blogList []models.Blog
	o := orm.NewOrm()
	qs := o.QueryTable("Blog")
	//var blogList []orm.ParamsList
	num, err := qs.Filter("id__gte", 1).All(&blogList)
	if num > 0 && err == nil {
		success = true
	} else {
		respone = "4001"
	}
	if success == true {
		this.Data["json"] = map[string]interface{}{"status": 0, "desc": respone, "data": blogList}
	} else {
		this.Data["json"] = map[string]interface{}{"status": 0, "desc": respone}
	}
	this.Data["blogs"] = blogList
	this.Layout = "layout.tpl"
	this.TplName = "index.tpl"
	//this.ServeJSON()
}

func (this *BlogController) GetByIdAndDelete() {
	id, _ := this.GetInt("id")
	deleted, _ := this.GetInt("deleted")

	success := false
	respone := "0"
	var blog models.Blog

	if id > 0 {
		o := orm.NewOrm()
		//qs := o.QueryTable("Blog")
		cond := orm.NewCondition()
		cond = cond.And("Id", id)

		cond1 := orm.NewCondition()
		cond1 = cond1.And("Deleted", deleted)

		cond = cond.AndCond(cond1)

		err := o.QueryTable("Blog").SetCond(cond).One(&blog)
		if err == nil {
			success = true
		} else {
			respone = "4001"
		}
	} else {
		respone = "4000"
	}

	if success == true {
		this.Data["json"] = map[string]interface{}{"status": 0, "desc": respone, "data": blog}
	} else {
		this.Data["json"] = map[string]interface{}{"status": 0, "desc": respone}
	}
	this.ServeJSON()

}

func (this *BlogController) DeleteBlog() {
	//id, _ := this.Ctx.Input.Params(":id")
	//fmt.Println("=================", id)
	id, _ := this.GetInt("id")
	success := false
	respone := "0"
	if id > 0 {
		o := orm.NewOrm()
		qs := o.QueryTable("Blog")
		_, err := qs.Filter("Id", id).Delete()
		if err == nil {
			success = true
		} else {
			respone = "4001"
		}
	} else {
		respone = "4000"
	}

	if success == true {
		this.Data["json"] = map[string]interface{}{"status": 0, "desc": respone}
	} else {
		this.Data["json"] = map[string]interface{}{"status": 1, "desc": respone}
	}
	this.ServeJSON()
}

func (this *BlogController) AddBlog() {
	title := this.GetString("title")
	content := this.GetString("content")
	fmt.Println("title=========", title)
	fmt.Println("content=========", content)

	success := false
	respone := "0"
	var blog models.Blog
	if len(title) > 0 && len(content) > 0 {
		o := orm.NewOrm()
		blog.Title = title
		blog.Content = content
		t := time.Now()
		blog.Created = t
		blog.Deleted = 0
		_, err := o.Insert(&blog)
		if err == nil {
			success = true

		} else {
			respone = "4001"
		}
	} else {
		respone = "4000"
	}
	if success == true {
		this.Data["json"] = map[string]interface{}{"status": 0, "desc": respone}
	} else {
		this.Data["json"] = map[string]interface{}{"status": 1, "desc": respone}
	}
	this.ServeJSON()
}

func (this *BlogController) UpdateBlog() {
	id, _ := this.GetInt("id")
	title := this.GetString("title")
	success := false
	respone := "0"
	param := orm.Params{"title": title}
	if id > 0 && len(title) > 0 {
		o := orm.NewOrm()
		qs := o.QueryTable("Blog")
		_, err := qs.Filter("id", id).Update(param)
		if err == nil {
			success = true
		} else {
			respone = "4001"
		}
	} else {
		respone = "4000"
	}
	if success == true {
		this.Data["json"] = map[string]interface{}{"status": 0, "desc": respone}
	} else {
		this.Data["json"] = map[string]interface{}{"status": 1, "desc": respone}
	}
	this.ServeJSON()
}
