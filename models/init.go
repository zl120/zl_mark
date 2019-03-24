package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql",orm.DRMySQL)
	orm.RegisterDataBase("default","mysql","mysql:Huawei@123@tcp(116.62.60.254:3306)/zl_test")
	orm.RegisterModel(&UserAccount{},&UserInfo{})
	orm.RunSyncdb("default",false,true)
}
