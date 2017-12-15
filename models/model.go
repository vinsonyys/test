package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"net/url"
	"time"
)

type Blog struct {
	Id      int `PK`
	Title   string
	Content string
	Created time.Time
	Deleted int
}

func init() {
	orm.RegisterModel(new(Blog))

	orm.RegisterDriver("mysql", orm.DRMySQL)
	connectStr := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&loc=%s", beego.AppConfig.String("mysqluser"), beego.AppConfig.String("mysqlpass"),
		beego.AppConfig.String("mysqlurls"), beego.AppConfig.String("mysqldb"), url.QueryEscape("Asia/Shanghai"))

	fmt.Println(connectStr)
	orm.RegisterDataBase("default", "mysql", connectStr) //创建数据库链接可创建链接池
	orm.Debug = true

}

// //获得数据库连接
// func GetLink() beedb.Model {
// 	//db, err := sql.Open("mymysql", "blog/astaxie/123456")
// 	db, err := sql.Open("mysql", "root:@/test?charset=utf8")
// 	if err != nil {
// 		panic(err)
// 	}
// 	orm := beedb.New(db)
// 	return orm
// }

// func GetAll() (blogs []Blog) {
// 	db := GetLink()
// 	db.FindAll(&blogs)
// 	return
// }

// func GetBlog(id int) (blog Blog) {
// 	db := GetLink()
// 	db.Where("id=?", id).Find(&blog)
// 	return
// }

// func SaveBlog(blog Blog) (bg Blog) {
// 	db := GetLink()
// 	db.Save(&blog)
// 	return bg
// }

// func DelBlog(blog Blog) {
// 	db := GetLink()
// 	db.Delete(&blog)
// 	return
// }
