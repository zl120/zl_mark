package controllers

import (
	"encoding/json"
	"zl_mark/models/vo"
	"zl_mark/service"
	"zl_mark/utils"
)

type UserInfoController struct {
	MainController
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

func (c *UserInfoController) CreateUserInfo() {
	var data map[string]interface{}
	if err := c.isNetAttack(); err!=nil{
		data["code"] = "failed"
		data["error"] = err.Error()
		c.Data["json"],err = utils.DataToMap(data)
		if err != nil {
			c.Data["json"] = "unknown error"
		}
		return
	}
	var userSrc *service.UserInfoService = &service.UserInfoService{}
	var user *vo.UserInfo = &vo.UserInfo{}
	input:=c.Ctx.Input.RequestBody
	if err:=json.Unmarshal(input,&user);err!=nil{
		data["code"] = "failed"
		data["error"] = err.Error()
		c.Data["json"],err = utils.DataToMap(data)
		if err != nil {
			c.Data["json"] = "unknown error"
		}
		return
	}
	userSrc.Name = user.Name
	userSrc.Age = user.Age
	userSrc.Tel = user.Tel
	userSrc.IdCard = user.IdCard
	userSrc.Gender = user.Gender
	id,err:=userSrc.CreateUserInfo()
	if err != nil {
		data["code"] = "failed"
		data["error"] = err.Error()
		c.Data["json"],err = utils.DataToMap(data)
		if err != nil {
			c.Data["json"] = "unknown error"
		}
		return
	}
	c.Data["json"] = id
}
