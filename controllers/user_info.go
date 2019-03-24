package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"zl_mark/models/vo"
	"zl_mark/service"
)

type UserInfoController struct {
	beego.Controller
	Id          int
	Name        string
	Age         int
	IdCard      string
	Gender      string
	Token       string
	Status      int
	Description string
	Tel         string
}

//judge that request whether a net attack?
func (c *UserInfoController) isNetAttack() error {
	return nil
}

func (c *UserInfoController) CreateUserAccount() {
	if err := c.isNetAttack(); err!=nil{
		c.Data["json"] = "error"
		return
	}
	var userSrc *service.UserInfoService = &service.UserInfoService{}
	var user *vo.UserInfo = &vo.UserInfo{}
	input:=c.Ctx.Input.RequestBody
	if err:=json.Unmarshal(input,&user);err!=nil{
		c.Data["json"] = "error"
		return
	}
	userSrc.Name = user.Name
	userSrc.Age = user.Age
	userSrc.Tel = user.Tel
	userSrc.IdCard = user.IdCard
	userSrc.Gender = user.Gender
	id,err:=userSrc.CreateUserInfo()
	if err != nil {
		c.Data["json"] = "error"
		return
	}
	c.Data["json"] = id
}
