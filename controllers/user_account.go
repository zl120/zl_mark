package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"strconv"
	"zl_mark/models/vo"
	"zl_mark/service"
)

type UserAccountController struct {
	beego.Controller
	Id          int
	Name        string
	Pwd         string
	UserInfoId  int
	Token       string
	Status      int
	Description string
	Role        string

}

func (c *UserAccountController) CreateUserAccount()  {
	input:=c.Ctx.Input.RequestBody
	var user  *vo.UserAccount = &vo.UserAccount{}
	if err:=json.Unmarshal(input,user);err!=nil{
		c.Data["error"] = ""
		c.Data["json"] = ""
		return
	}
	var userSrc *service.UserAccountService = &service.UserAccountService{}
	userSrc.Name = user.Name
	userSrc.Pwd = user.Pwd
	id,err:=userSrc.CreateUserAccount();
	if err!=nil{
		c.Data["error"] = ""
		c.Data["json"] = ""
		return
	}
	c.Data["error"] = ""
	c.Data["json"] = "id"+strconv.Itoa(id)
	return
}
