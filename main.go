package main

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	_ "zl_mark/routers"
)

func main() {
	beego.Run()
}

