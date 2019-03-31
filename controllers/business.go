package controllers

import "html/template"

type BusinessController struct {
	MainController

}

func (c *BusinessController) GetHome()  {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "test/home.html"
}

func (c *BusinessController) Login(){
	c.Data["xsrfdata"]=template.HTML(c.XSRFFormHTML())
	c.TplName = "test/login.html"
}

func (c *BusinessController) Register(){
	c.Data["xsrfdata"]=template.HTML(c.XSRFFormHTML())
	c.TplName = "test/register.html"
}
func (c *BusinessController) ForgetPwd(){
	c.TplName = "test/forget-pwd.html"
}