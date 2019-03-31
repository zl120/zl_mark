package controllers

import (
	"fmt"
	"strconv"
	"zl_mark/service"
	"zl_mark/utils"
)

type UserAccountController struct {
	MainController
	Id          int
	Name        string
	Pwd         string
	UserInfoId  int
	Token       string
	Status      int
	Description string
	Role        string
}

func (c *UserAccountController) Register() {
	var data map[string]interface{} = make(map[string]interface{})
	var err error
	defer func() {
		if err != nil {
			data["code"] = "failed"
			data["error"] = err.Error()
		} else {
			data["code"] = "success"
		}
		c.Data["json"], err = utils.DataToMap(data)
		if err != nil {
			c.Data["json"] = "unknown error"
		}
		c.ServeJSON()
	}()
	name := c.Ctx.Request.FormValue("name")
	pwd := c.Ctx.Request.FormValue("pwd")
	var userSrc *service.UserAccountService = &service.UserAccountService{Name:name,Pwd:pwd}
	var id int
	id, err = userSrc.CreateUserAccount();
	if err != nil {
		return
	}
	data["id"] = strconv.Itoa(id)
}

func (c *UserAccountController) Login() {
	var data map[string]interface{} = make(map[string]interface{})
	var err error
	defer func() {
		if err != nil {
			data["code"] = "failed"
			data["error"] = err.Error()
			c.Data["json"], err = utils.DataToMap(data)
			if err != nil {
				c.Data["json"] = "unknown error"
			}
			fmt.Println("json")
			c.ServeJSON()
		}
	}()
	name := c.Ctx.Request.FormValue("name")
	pwd := c.Ctx.Request.FormValue("pwd")
	var userSrv *service.UserAccountService = &service.UserAccountService{Name: name, Pwd: pwd}
	if err = userSrv.Login(); err != nil {
		return
	}
	c.TplName = "test/home.html"
}
