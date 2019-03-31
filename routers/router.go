package routers

import (
	"zl_mark/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.BusinessController{},"get:Login")
	beego.Router("/register", &controllers.BusinessController{},"get:Register")
	beego.Router("/forget-pwd", &controllers.BusinessController{},"get:ForgetPwd")
	beego.Router("/delete", &controllers.BusinessController{},"get:Delete")
	beego.Router("/login", &controllers.UserAccountController{},"post:Login")
	beego.Router("/register", &controllers.UserAccountController{},"post:Register")
	beego.Router("/home", &controllers.BusinessController{},"get:GetHome")
	beego.Router("/forget-pwd", &controllers.MainController{})
}
